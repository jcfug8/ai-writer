package service

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

// Opts - Options for control service
type Opts struct {
	Addr string
}

// Service - Persist Service
type Service struct {
	Opts       *Opts
	router     *mux.Router
	httpServer *http.Server
	db         *sql.DB
}

// NewService - Takes Opts and returns an initialized control service
func NewService(opts *Opts) *Service {
	s := &Service{
		Opts:   opts,
		router: mux.NewRouter().StrictSlash(true),
	}

	// set up user routes
	log.Info("about to init user router")
	userRouter := s.router.PathPrefix("/user").Subrouter()
	userRouter.Path("/").Methods("GET").HandlerFunc(s.GetUser)
	// set up book routes
	log.Info("about to init book router")
	bookRouter := s.router.PathPrefix("/book").Subrouter()
	bookRouter.Path("/").Methods("GET").HandlerFunc(s.GetBook)
	bookRouter.Path("/").Methods("POST").HandlerFunc(s.CreateBook)

	// set up http server
	s.httpServer = &http.Server{
		Handler:      s.router,
		Addr:         opts.Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return s
}

// Serve - Start up the persist service
func (s *Service) Serve() error {
	log.Info("serving persist")

	return s.httpServer.ListenAndServe()
}

func (s *Service) GetUser(w http.ResponseWriter, r *http.Request) {
	log.Info("Geting User")
	w.Write([]byte("Get"))
}

func (s *Service) GetBook(w http.ResponseWriter, r *http.Request) {
	log.Info("Geting Book")
	w.Write([]byte("Get"))
}

func (s *Service) CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Info("Creating Book")
	w.Write([]byte("Create"))
}
