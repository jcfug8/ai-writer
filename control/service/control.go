package service

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

// Opts - Options for control service
type Opts struct {
	AIAddr      string
	PersistAddr string
	Addr        string
	AssetsDir   string
}

// Service - Control Service
type Service struct {
	Opts       *Opts
	router     *mux.Router
	fileServer http.Handler
	httpServer *http.Server
}

// NewService - Takes Opts and returns an initialized control service
func NewService(opts *Opts) *Service {
	s := &Service{
		Opts:       opts,
		router:     mux.NewRouter().StrictSlash(true),
		fileServer: http.FileServer(http.Dir(opts.AssetsDir)),
	}

	// set up api routes
	// api := s.router.PathPrefix("/api/").Subrouter()

	// set up static file routes
	s.router.PathPrefix("/").Handler(s.fileServer)
	s.router.Path("/").HandlerFunc(s.serveIndex)

	// set up http server
	s.httpServer = &http.Server{
		Handler:      s.router,
		Addr:         opts.Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return s
}

// Serve - Start up the control service
func (s *Service) Serve() error {
	log.Info("serving control")
	return s.httpServer.ListenAndServe()
}

func (s *Service) serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(s.Opts.AssetsDir, "index.html"))
	return
}
