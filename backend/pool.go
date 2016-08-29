package backend

import (
	"fmt"
	. "github.com/ruiaylin/flowage/mysql"
	"github.com/ruiaylin/flowage/utils/errors"
	"sync"
	"sync/atomic"
	"time"
)

//   mincached: initial number of idle connections in the pool
// 	     (0 means no connections are made at startup)
// 	 maxcached: maximum number of idle connections in the pool
// 	     (0 or None means unlimited pool size)
// 	 maxconnections: maximum number of connections generally allowed
// 	     (0 or None means an arbitrary number of connections)
// 	 blocking: determines behavior when exceeding the maximum
// 	     (if this is set to true, block and wait until the number of
// 	     connections decreases, otherwise an error will be reported)
// 	 ping: determines when the connection should be checked with ping()
// 	     (0 = None = never, 1 = default = whenever fetched from the pool,
// 	     2 = when a cursor is created, 4 = when a query is executed,
// 	     7 = always, and all other bit combinations of these values)
// 	 args, kwargs: the parameters that shall be passed to the creator
// 	     function or the connection constructor of the DB-API 2 module

const (
	// pool state
	Up = iota
	Down
	ManualDown
	Unknown

	// pool
	PoolInitSize           = 8
	PoolMaxCacheSize       = 16
	PoolMaxSize            = 10
	CheckPeroid      int64 = 4
)

type Pool struct {
	// lock
	sync.RWMutex
	// connection information
	addr     string
	user     string
	password string
	database string // as schema
	state    int32
	// pool control
	unlimited  bool //if unlimit is True , the pool will create new conn when the max reach
	minCached  int32
	maxCached  int32
	maxConnNum int32
	connNum    int32
	// for query timeout
	blockChannel    chan<- time.Duration
	blockChMux      sync.RWMutex
	usageTimeout    time.Duration
	usageTimeoutCh  chan<- string
	usageTimeoutMux sync.RWMutex
	// connection containers
	// cached connection
	cacheQueue chan *Conn
	// idle connection
	idleQueue chan *Conn
	// check connecttion
	checkConn *Conn
}

func OpenPool(addr string, user string, password string, dbName string, minCached int32, maxCached int32, maxConnNum int32, unlimit bool) (*Pool, error) {
	var err error
	p := new(Pool)
	p.addr = addr
	p.user = user
	p.password = password
	p.database = dbName
	// 8 , 10> 10,  4> 8
	p.unlimited = unlimit
	if minCached > 0 {
		p.minCached = minCached
	} else {
		p.minCached = 0
	}
	// 0 unlimit
	if maxCached > 0 {
		if maxCached < minCached {
			maxCached = minCached
		}
		p.maxCached = maxCached
	} else {
		p.maxCached = PoolMaxCacheSize
	}

	// 0 unlimited , just for max pool size
	if maxConnNum > 0 {
		if maxConnNum < p.maxCached {
			maxConnNum = p.maxCached
		}
		p.maxConnNum = maxConnNum
	} else {
		maxConnNum = PoolMaxSize
		p.unlimited = true
	}

	// check connection
	p.checkConn, err = p.newConn()
	if err != nil {
		p.Close()
		return nil, err
	}
	// conns container
	p.cacheQueue = make(chan *Conn, p.maxCached)
	p.idleQueue = make(chan *Conn, p.maxConnNum)
	atomic.StoreInt32(&(p.state), Unknown)

	// init the pool
	var i int32
	for i = 0; i < maxConnNum; i++ {
		if i < p.minCached {
			conn, err := p.newConn()
			if err != nil {
				p.Close()
				return nil, err
			}
			conn.lastPing = time.Now().Unix()
			p.cacheQueue <- conn
		} else {
			conn := new(Conn)
			p.idleQueue <- conn
		}
	}
	p.connNum = maxConnNum
	return p, nil
}

func (p *Pool) Close() error {
	p.Lock()
	cacheChannel := p.cacheQueue
	p.cacheQueue = nil
	p.Unlock()
	if cacheChannel == nil {
		return nil
	}

	close(cacheChannel)
	for conn := range cacheChannel {
		p.closeConn(conn)
	}

	return nil
}

func (p *Pool) newConn() (*Conn, error) {
	co := new(Conn)
	if err := co.Connect(p.addr, p.user, p.password, p.database); err != nil {
		return nil, err
	}
	return co, nil
}

func (p *Pool) closeConn(co *Conn) error {
	if co != nil {
		co.Close()
		p.RLock()
		conns := p.idleQueue
		p.RUnlock()
		if conns != nil {
			select {
			case conns <- co:
				return nil
			default:
				return nil
			}
		}
	}
	return nil
}

func (p *Pool) State() string {
	var state string
	switch p.state {
	case Up:
		state = "up"
	case Down, ManualDown:
		state = "down"
	case Unknown:
		state = "unknow"
	}
	return state
}

// check the pool's status
func (p *Pool) Ping() error {
	var err error
	if p.checkConn == nil {
		p.checkConn, err = p.newConn()
		if err != nil {
			p.closeConn(p.checkConn)
			p.checkConn = nil
			return err
		}
	}
	err = p.checkConn.Ping()
	if err != nil {
		p.closeConn(p.checkConn)
		p.checkConn = nil
		return err
	}
	return nil
}

// handle reuse a connection
func (p *Pool) tryReuse(co *Conn) error {
	if co.IsInTransaction() {
		//we can not reuse a connection in transaction status
		if err := co.Rollback(); err != nil {
			return err
		}
	}

	if !co.IsAutoCommit() {
		//we can not  reuse a connection not in autocomit
		if _, err := co.exec("set autocommit = 1"); err != nil {
			return err
		}
	}

	//connection may be set names early
	//we must use default utf8
	if co.GetCharset() != DEFAULT_CHARSET {
		if err := co.SetCharset(DEFAULT_CHARSET); err != nil {
			return err
		}
	}

	return nil
}

func (p *Pool) Addr() string {
	return p.addr
}

func (p *Pool) String() string {
	return fmt.Sprintf("%s:%s@%s/%s",
		p.user, p.password, p.addr, p.database)
}

func (p *Pool) GetConnFromCache(cacheQueue chan *Conn) *Conn {
	var co *Conn
	var err error
	for 0 < len(cacheQueue) {
		co = <-cacheQueue
		if co != nil && CheckPeroid < time.Now().Unix()-co.lastPing {
			err = co.Ping()
			if err != nil {
				p.closeConn(co)
				co = nil
			}
		}
		if co != nil {
			break
		}
	}
	return co
}

func (p *Pool) GetConnFromIdle(cacheQueue, idleQueue chan *Conn) (*Conn, error) {
	var co *Conn
	var err error
	select {
	case co = <-idleQueue:
		err = co.Connect(p.addr, p.user, p.password, p.database)
		if err != nil {
			p.closeConn(co)
			return nil, err
		}
		return co, nil
	case co = <-cacheQueue:
		if co == nil {
			return nil, errors.ErrConnIsNil
		}
		if co != nil && CheckPeroid < time.Now().Unix()-co.lastPing {
			err = co.Ping()
			if err != nil {
				p.closeConn(co)
				return nil, errors.ErrBadConn
			}
		}
	default:
		// if maxConnNum = 0 , unlimited the maxsize
		// or the unilimit parameter set true
		if p.unlimited {
			co, err = p.newConn()
			if err != nil {
				p.Close()
				return nil, err
			}
			co.lastPing = time.Now().Unix()
			p.RLock()
			p.connNum++
			p.RUnlock()
		} else {
			return nil, errors.ErrNoAvailConn
		}
	}
	return co, nil
}

func (p *Pool) CachedConnCount() int32 {
	p.RLock()
	defer p.RUnlock()
	return int32(len(p.cacheQueue))
}

func (p *Pool) IdelConnCount() int {
	p.RLock()
	defer p.RUnlock()
	return len(p.idleQueue)
}

func (p *Pool) GetConnNum() int32 {
	return atomic.LoadInt32(&p.connNum)
}

func (p *Pool) GetMaxCached() int32 {
	return atomic.LoadInt32(&p.maxCached)
}
func (p *Pool) GetMinCached() int32 {
	return atomic.LoadInt32(&p.minCached)
}

func (p *Pool) GetMaxConnNum() int32 {
	return atomic.LoadInt32(&p.maxConnNum)
}

func (p *Pool) PushConn(co *Conn, err error) {
	if p.CachedConnCount() < p.maxCached {
		if co == nil {
			return
		}
		p.RLock()
		conns := p.cacheQueue
		p.RUnlock()
		if conns == nil {
			co.Close()
			return
		}
		if err != nil {
			p.closeConn(co)
			return
		}
		co.lastPing = time.Now().Unix()
		select {
		case conns <- co:
			return
		default:
			p.closeConn(co)
			return
		}
	} else {
		p.closeConn(co)
		return
	}
}

func (p *Pool) PopConn() (*Conn, error) {
	var co *Conn
	var err error
	// get the queue
	p.RLock()
	cacheQueue := p.cacheQueue
	idleQueue := p.idleQueue
	p.RUnlock()

	if cacheQueue == nil || idleQueue == nil {
		return nil, errors.ErrDatabaseClose
	}
	co = p.GetConnFromCache(cacheQueue)
	if co == nil {
		co, err = p.GetConnFromIdle(cacheQueue, idleQueue)
		if err != nil {
			fmt.Printf("PopConn error : %s \n", err)
			return nil, err
		}
	}

	err = p.tryReuse(co)
	if err != nil {
		p.closeConn(co)
		return nil, err
	}

	return co, nil
}

// get one connection from pool
func (p *Pool) GetConn() (*PoolConn, error) {
	c, err := p.PopConn()
	if err != nil {
		fmt.Printf("GetConn error : %s \n", err)
		return nil, err
	}
	return &PoolConn{c, p}, nil
}

// track the relation between pool and the connection
type PoolConn struct {
	*Conn
	pool *Pool
}

// close connection to the pool
func (pc *PoolConn) Close() {
	if pc != nil && pc.Conn != nil {
		if pc.Conn.pkgErr != nil {
			pc.pool.closeConn(pc.Conn)
		} else {
			pc.pool.PushConn(pc.Conn, nil)
		}
		pc.Conn = nil
	}
}
