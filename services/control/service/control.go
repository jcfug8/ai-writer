package service

import (
	"context"
	"encoding/json"
	"net/http"
	"path/filepath"
	"time"

	pb "github.com/jcfug8/ai-writer/protos"

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
	opts          *Opts
	router        *mux.Router
	fileServer    http.Handler
	httpServer    *http.Server
	persistClient pb.PersistClient
}

// NewService - Takes Opts and returns an initialized control service
func NewService(p pb.PersistClient, opts *Opts) *Service {
	s := &Service{
		opts:          opts,
		router:        mux.NewRouter().StrictSlash(true),
		fileServer:    http.FileServer(http.Dir(opts.AssetsDir)),
		persistClient: p,
	}

	// set up api routes
	s.router.Methods("GET").Path("/api/user").HandlerFunc(s.getUser)
	s.router.Methods("GET").Path("/api/book").HandlerFunc(s.getBook)
	s.router.Methods("POST").Path("/api/book").HandlerFunc(s.createBook)

	// set up static file routes
	s.router.Methods("GET").PathPrefix("/assets/").Handler(s.fileServer)
	s.router.Methods("GET").Path("/").HandlerFunc(s.serveIndex)

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
	log.Infof("serving control at %s", s.opts.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Service) serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(s.opts.AssetsDir, "index.html"))
	return
}

func (s *Service) getUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	res, err := s.persistClient.GetUser(ctx, &pb.GetUserRequest{})
	cancel()
	if err != nil {
		e, _ := json.Marshal(err)
		w.Write(e)
		return
	}
	br, _ := json.Marshal(res)
	w.Write(br)
	return
}

func (s *Service) getBook(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	res, err := s.persistClient.GetBook(ctx, &pb.GetBookRequest{})
	cancel()
	if err != nil {
		e, _ := json.Marshal(err)
		w.Write(e)
		return
	}
	br, _ := json.Marshal(res)
	w.Write(br)
	return
}

func (s *Service) createBook(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	res, err := s.persistClient.CreateBook(ctx, &pb.CreateBookRequest{})
	cancel()
	if err != nil {
		e, _ := json.Marshal(err)
		w.Write(e)
		return
	}
	br, _ := json.Marshal(res)
	w.Write(br)
	return
}
