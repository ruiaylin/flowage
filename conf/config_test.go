package conf

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	var testConfigData = []byte(
		`
    {
    "addr":"127.0.0.1:4000",
    "user":"root",
    "password":"",
    "StatusAddr":":10080"
    "groups":[
            {
                "name": "trade00",
                "RWSplit":True,
                "failover":"",
                "master": {
                        "name": "node1",
                        "role":"master",
                        "addr":"127.0.0.1",
                        "port":3306 ,
                        "password":"",
                        "mincached":16,
                        "maxcached":64,
                        "maxConnNum":128,
                        "unlimit":True
                    }, 
                "slave": {
                        "name": "node2",
                        "role":"slave",
                        "addr":"127.0.0.1",
                        "port":3307 ,
                        "password":"",
                        "mincached":16,
                        "maxcached":32,
                        "maxConnNum":64,
                        "unlimit":false
                    },
            },
            {
                "name": "trade01",
                "RWSplit":True,
                "failover":"",
                "master": {
                        "name": "node3",
                        "role":"master",
                        "addr":"127.0.0.2",
                        "port":3306 ,
                        "password":"",
                        "mincached":16,
                        "maxcached":64,
                        "maxConnNum":128,
                        "unlimit":True
                    }, 
                "slave": {
                        "name": "node4",
                        "role":"slave",
                        "addr":"127.0.0.2",
                        "port":3307 ,
                        "password":"",
                        "mincached":16,
                        "maxcached":32,
                        "maxConnNum":64,
                        "unlimit":false
                    }
            }
        ],
    "schema": [{
            "db": "flowage00",
            "node":  "trade00"
            },{
                "db": "flowage01",
                "node":  "trade00"
            },{
                "db": "flowage02",
                "node":  "trade00"
            },{
                "db": "flowage03",
                "node":  "trade00"
            },{
                "db": "flowage04",
                "node":  "trade00"
            },{
                "db": "flowage05",
                "node":  "trade01"
            },{
                "db": "flowage06",
                "node":  "trade01"
            },{
                "db": "flowage07",
                "node":  "trade01"
            },{
                "db": "flowage08",
                "node":  "trade01"
            },{
                "db": "flowage09",
                "node":  "trade01"
            }
        ]
    "tables":[{
            "table"   : "trade",
            "pkey"    : "tradeno",
            "type"    : "char",
            "method"  : "hash",
            "partitions":
               [
                   { "suffix" : "_0", "db": "flowage00" },
                   { "suffix" : "_1", "db": "flowage01" },
                   { "suffix" : "_2", "db": "flowage02" },
                   { "suffix" : "_3", "db": "flowage03" },
                   { "suffix" : "_4", "db": "flowage04" },
                   { "suffix" : "_5", "db": "flowage05" },
                   { "suffix" : "_6", "db": "flowage06" },
                   { "suffix" : "_7", "db": "flowage07" },
                   { "suffix" : "_8", "db": "flowage08" },
                   { "suffix" : "_9", "db": "flowage09" } 
               ]
            },{
                "table"   : "trade_order",
                "pkey"    : "tradeno",
                "type"    : "char",
                "method"  : "hash",
                "partitions":
                   [
                       { "suffix" : "_0", "db": "flowage00" },
                       { "suffix" : "_1", "db": "flowage01" },
                       { "suffix" : "_2", "db": "flowage02" },
                       { "suffix" : "_3", "db": "flowage03" },
                       { "suffix" : "_4", "db": "flowage04" },
                       { "suffix" : "_5", "db": "flowage05" },
                       { "suffix" : "_6", "db": "flowage06" },
                       { "suffix" : "_7", "db": "flowage07" },
                       { "suffix" : "_8", "db": "flowage08" },
                       { "suffix" : "_9", "db": "flowage09" } 
                   ]
            },{
                "table"   : "order_ext",
                "pkey"    : "tradeno",
                "type"    : "char",
                "method"  : "hash",
                "partitions":
                   [
                       { "suffix" : "_0", "db": "flowage00" },
                       { "suffix" : "_1", "db": "flowage01" },
                       { "suffix" : "_2", "db": "flowage02" },
                       { "suffix" : "_3", "db": "flowage03" },
                       { "suffix" : "_4", "db": "flowage04" },
                       { "suffix" : "_5", "db": "flowage05" },
                       { "suffix" : "_6", "db": "flowage06" },
                       { "suffix" : "_7", "db": "flowage07" },
                       { "suffix" : "_8", "db": "flowage08" },
                       { "suffix" : "_9", "db": "flowage09" } 
                   ],
                subpkey : create_time
                subtype : timestamp
                submethod : range
                subpartitions : [   { "suffix" : "_16Q1", "value" : "2016/04/01" },
                                    { "suffix" : "_16Q2", "value" : "2016/07/01" },
                                    { "suffix" : "_16Q3", "value" : "2016/10/01" },
                                    { "suffix" : "_16Q4", "value" : "2017/01/01" }]

            }
        ]

}
`)

	cfg, err := ParseConfigData(testConfigData)
	if err != nil {
		t.Fatal(err)
	}

	if len(cfg.Groups) != 3 {
		t.Fatal(len(cfg.Groups))
	}

	if len(cfg.Schemas) != 1 {
		t.Fatal(len(cfg.Schemas))
	}

	testNode := NodeConfig{
		Name:             "node1",
		DownAfterNoAlive: 300,
		IdleConns:        16,
		RWSplit:          true,

		User:     "root",
		Password: "",

		Master: "127.0.0.1:3306",
		Slave:  "127.0.0.1:4306",
	}

	if !reflect.DeepEqual(cfg.Nodes[0], testNode) {
		fmt.Printf("%v\n", cfg.Nodes[0])
		t.Fatal("node1 must equal")
	}

	testNode_2 := NodeConfig{
		Name:   "node2",
		User:   "root",
		Master: "127.0.0.1:3307",
	}

	if !reflect.DeepEqual(cfg.Nodes[1], testNode_2) {
		t.Fatal("node2 must equal")
	}

	testShard_1 := ShardConfig{
		Table: "flowage_test_shard_hash",
		Key:   "id",
		Nodes: []string{"node1", "node2", "node3"},
		Type:  "hash",
	}
	if !reflect.DeepEqual(cfg.Schemas[0].RulesConifg.ShardRule[0], testShard_1) {
		t.Fatal("ShardConfig0 must equal")
	}

	testShard_2 := ShardConfig{
		Table: "flowage_test_shard_range",
		Key:   "id",
		Nodes: []string{"node2", "node3"},
		Type:  "range",
		Range: "-10000-",
	}
	if !reflect.DeepEqual(cfg.Schemas[0].RulesConifg.ShardRule[1], testShard_2) {
		t.Fatal("ShardConfig1 must equal")
	}

	if 2 != len(cfg.Schemas[0].RulesConifg.ShardRule) {
		t.Fatal("ShardRule must 2")
	}

	testRules := RulesConfig{
		Default:   "node1",
		ShardRule: []ShardConfig{testShard_1, testShard_2},
	}
	if !reflect.DeepEqual(cfg.Schemas[0].RulesConifg, testRules) {
		t.Fatal("RulesConfig must equal")
	}

	testSchema := SchemaConfig{
		DB:          "flowage",
		Nodes:       []string{"node1", "node2", "node3"},
		RulesConifg: testRules,
	}

	if !reflect.DeepEqual(cfg.Schemas[0], testSchema) {
		t.Fatal("schema must equal")
	}

	if cfg.LogLevel != "error" || cfg.User != "root" || cfg.Password != "" || cfg.Addr != "127.0.0.1:4000" {
		t.Fatal("Top Config not equal.")
	}
}
