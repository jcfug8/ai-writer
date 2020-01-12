package main

import (
	"flag"
	"path/filepath"

	"github.com/jcfug8/ai-writer/control/service"

	log "github.com/sirupsen/logrus"
)

var PersistAddr string
var AIAddr string
var Addr string
var AssetsDir string

func init() {
	flag.StringVar(&PersistAddr, "persist", "", "Address Persist service will be reached at")
	flag.StringVar(&AIAddr, "ai", "", "Address AI service will be reached at")
	flag.StringVar(&Addr, "a", "127.0.0.1:8080", "Address Control service will be served at")

	assetsDir, err := filepath.Abs("../../client")
	if err != nil {
		log.Panicf("unable to create default assests dir path: %s", err)
	}
	flag.StringVar(&AssetsDir, "assets", assetsDir, "Static files will be served from")
}

func main() {
	flag.Parse()
	log.WithFields(log.Fields{
		"persist": PersistAddr,
		"ai":      AIAddr,
		"a":       Addr,
		"assets":  AssetsDir,
	}).Info("parsed flags")

	control := service.NewService(&service.Opts{
		AIAddr:      AIAddr,
		PersistAddr: PersistAddr,
		Addr:        Addr,
		AssetsDir:   AssetsDir,
	})

	log.Info("serving control")
	if err := control.Serve(); err != nil {
		log.Errorf("control http server closed in error: %s", err)
	}
}
