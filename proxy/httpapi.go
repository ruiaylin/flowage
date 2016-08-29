package proxy

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ruiaylin/flowage/utils/glog/log"
	"net/http"
	"sync"
	// For pprof
)

var once sync.Once

const defaultStatusAddr = ":10080"

func (s *Server) startStatusHTTP() {
	once.Do(func() {
		go func() {
			// handle for status
			http.HandleFunc("/status", func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				s := status{TPS: 123, Connections: 120, Version: "5.6.25"}
				js, err := json.Marshal(s)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Error("Encode json error %s", err.Error())
				} else {
					w.Write(js)
				}
			})
			// handle for schema
			http.HandleFunc("/nodes", func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				log.Info("s.nodes : %s ", s.nodes)
				js, err := json.Marshal(s.nodes)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Error("Encode json error %s", err.Error())
				} else {
					w.Write(js)
				}
			})
			// handle for schema
			http.HandleFunc("/schema", func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				s := status{TPS: 123, Connections: 120, Version: s.schemas["flowage"].db}
				js, err := json.Marshal(s)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Error("Encode json error %s", err.Error())
				} else {
					w.Write(js)
				}

			})
			// HTTP path for prometheus.
			http.Handle("/metric", prometheus.Handler())
			addr := s.cfg.StatusAddr
			if len(addr) == 0 {
				addr = defaultStatusAddr
			}

			log.Info("Listening on %v for status and metrics report.", addr)
			err := http.ListenAndServe(addr, nil)
			if err != nil {
				log.Error("Encode json error %s", err.Error())
			}
		}()
	})
}

// Proxy status
type status struct {
	TPS         int64  `json:"tps"`
	Connections int    `json:"connections"`
	Version     string `json:"version"`
}

// Server error codes.
const (
	codeUnknownFieldType  = 1
	codeInvalidPayloadLen = 2
	codeInvalidSequence   = 3
	codeInvalidType       = 4
)
