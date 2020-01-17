package main

import (
	"flag"

	"github.com/jcfug8/ai-writer/services/persist/service"

	log "github.com/sirupsen/logrus"
)

var addr string

func init() {
	flag.StringVar(&addr, "a", "127.0.0.1:50051", "Address Persist service will be served at")
}

func main() {
	flag.Parse()
	log.WithFields(log.Fields{
		"a": addr,
	}).Info("parsed flags")

	persist := service.NewService(&service.Opts{
		Addr: addr,
	})

	log.Info("about to serve persist")
	if err := persist.Serve(); err != nil {
		log.Errorf("control http server closed in error: %s", err)
	}
}
