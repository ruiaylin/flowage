package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Addr       string         `json:"addr"`
	User       string         `json:"user"`
	Password   string         `json:"password"`
	LogLevel   string         `json:"log_level"`
	Groups     []GroupConfig  `json:"groups"`
	Schemas    []SchemaConfig `json:"schemas"`
	Tables     []ShardConfig  `json:"tables"`
	StatusAddr string         `json:"status_addr"`
}

// config
type GroupConfig struct {
	Name     string     `json:"name"` // like the schema
	RWSplit  bool       `json:"rw_split"`
	Failover NodeConfig `json:"failover"`
	Master   NodeConfig `json:"master"`
	Slave    NodeConfig `json:"slave"`
}

// group contain node
type NodeConfig struct {
	Name       string `json:"name"`
	Role       string `json:"role"`
	Addr       string `json:"addr"`
	Port       int32  `json:"port"`
	User       string `json:"user"`
	Password   string `json:"password"`
	MinCached  int32  `json:"mincached"`
	MaxCached  int32  `json:"maxcached"`
	MaxConnNum int32  `json:"maxconnnum"`
	Unlimit    bool   `json:"unlimit"`
}

type SchemaConfig struct {
	DB    string `json:"db"`
	Group string `json:"group"`
}

// Table map to group
type TableConfig struct {
	Suffix string `json:"suffix"`
	Schema string `json:"db"`
	value  string `json:vallue`
}

type RulesConfig struct {
	Default   string        `json:"default"`
	ShardRule []ShardConfig `json:"shard"`
}

type ShardConfig struct {
	Table         string        `json:"table"`
	PKey          string        `json:"pkey"`
	Type          string        `json:"type"`
	Method        string        `json:"method"`
	Partitions    []TableConfig `json:"partitions"`
	SubPkey       string        `json:"subpkey"`
	SubType       string        `json:"subtype"`
	SubMethod     string        `json:"submethod"`
	Subpartitions []TableConfig `json:"subpartitions"`
}

// for node
func (gc *GroupConfig) ToString() (string, error) {
	bstr, err := json.Marshal(gc)
	if err != nil {
		fmt.Println("ToString , err:  ", err)
		return " ", err
	}
	return string(bstr), err
}

// for node
func (nc *NodeConfig) ToString() (string, error) {
	bstr, err := json.Marshal(nc)
	if err != nil {
		fmt.Println("ToString , err:  ", err)
		return " ", err
	}
	return string(bstr), err
}

// for schema
func (sc *SchemaConfig) ToString() (string, error) {
	bstr, err := json.Marshal(sc)
	if err != nil {
		fmt.Println("ToString , err:  ", err)
		return " ", err
	}
	return string(bstr), err
}

// for shard
func (sc *ShardConfig) ToString() (string, error) {
	bstr, err := json.Marshal(sc)
	if err != nil {
		fmt.Println("ToString , err:  ", err)
		return " ", err
	}
	return string(bstr), err
}

// for table
func (tc *TableConfig) ToString() (string, error) {
	bstr, err := json.Marshal(tc)
	if err != nil {
		fmt.Println("ToString , err:  ", err)
		return " ", err
	}
	return string(bstr), err
}

func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := json.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return ParseConfigData(data)
}
