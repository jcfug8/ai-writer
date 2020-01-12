package main

import (
	"flag"

	"github.com/jcfug8/ai-writer/persist/service"

	log "github.com/sirupsen/logrus"
)

var PersistAddr string
var AIAddr string
var Addr string
var AssetsDir string

func init() {
	flag.StringVar(&Addr, "a", "127.0.0.1:8081", "Address Persist service will be served at")
}

func main() {
	flag.Parse()
	log.WithFields(log.Fields{
		"a": Addr,
	}).Info("parsed flags")

	persist := service.NewService(&service.Opts{
		Addr: Addr,
	})

	log.Info("about to serve persist")
	if err := persist.Serve(); err != nil {
		log.Errorf("control http server closed in error: %s", err)
	}
}
