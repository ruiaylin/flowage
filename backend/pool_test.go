package backend

import (
	"fmt"
	"testing"
)

func Test_Pooling(t *testing.T) {
	p, err := backend.OpenPool("127.0.0.1:3306", "flowage", "flowage", "flowage", 2, 0, 10, true)
	fmt.Printf(" pool : %s \n", p)
	if err != nil {
		fmt.Printf(" pool : %s \n", err)
		return
	}
	fmt.Printf(" dbns              : %s \n", p.String())
	fmt.Printf(" CachedConnCount   : %s \n", p.CachedConnCount())
	fmt.Printf(" IdelConnCount     : %s \n", p.IdelConnCount())
	fmt.Printf(" minCached         : %s \n", p.GetMinCached())
	fmt.Printf(" maxCached         : %s \n", p.GetMaxCached())
	fmt.Printf(" maxConnNum        : %s \n", p.GetMaxConnNum())
	var connlist [32]*backend.PoolConn
	x := 16
	i := 0
	for i < x {
		connlist[i], err = p.GetConn()
		if err != nil {
			fmt.Printf(" %v error  : %s \n", i, err)
		}
		i = i + 1
	}
	co1, err := p.GetConn()
	if err != nil {
		fmt.Printf(" %v error  : %s \n", i+1, err)
	}
	co1.Close()
	x = 16
	i = 0
	for i < x {
		connlist[i].Close()
		i = i + 1
	}
	p.Close()
}

func TestStmt_Delete(t *testing.T) {
	str := `delete from flowage_test_stmt`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Execute(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}
