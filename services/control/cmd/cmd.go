package main

import (
	"flag"
	"path/filepath"

	pb "github.com/jcfug8/ai-writer/protos"

	"github.com/jcfug8/ai-writer/services/control/service"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

var persistAddr string
var aiAddr string
var addr string
var assetsDir string

func init() {
	flag.StringVar(&persistAddr, "persist", "127.0.0.1:50051", "Address Persist service will be reached at")
	flag.StringVar(&aiAddr, "ai", "127.0.0.1:50051", "Address AI service will be reached at")
	flag.StringVar(&addr, "a", "127.0.0.1:8080", "Address Control service will be served at")

	_assetsDir, err := filepath.Abs("../../client")
	if err != nil {
		log.Panicf("unable to create default assests dir path: %s", err)
	}
	flag.StringVar(&assetsDir, "assets", _assetsDir, "Static files will be served from")
}

func main() {
	flag.Parse()
	log.WithFields(log.Fields{
		"persist": persistAddr,
		"ai":      aiAddr,
		"a":       addr,
		"assets":  assetsDir,
	}).Info("parsed flags")

	// Create Persist Client
	log.Infof("creating persist connection and client at %s", persistAddr)
	persistConn, err := grpc.Dial(persistAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to persist client at %s: %s", persistAddr, err)
	}
	persistClient := pb.NewPersistClient(persistConn)

	// Create AI Client
	log.Infof("creating ai connection and client at %s", aiAddr)
	aiConn, err := grpc.Dial(aiAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to ai client at %s: %s", aiAddr, err)
	}
	aiClient := pb.NewAIClient(aiConn)

	control := service.NewService(persistClient, aiClient, &service.Opts{
		Addr:      addr,
		AssetsDir: assetsDir,
	})

	log.Info("about to serve control")
	if err := control.Serve(); err != nil {
		log.Errorf("control http server closed in error: %s", err)
	}
	persistConn.Close()
}
