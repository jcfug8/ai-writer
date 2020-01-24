package service

import (
	"context"
	"encoding/json"
	"net/http"
	"path/filepath"
	"time"

	"github.com/jcfug8/ai-writer/commons/errors"
	"github.com/jcfug8/ai-writer/commons/replies"
	pb "github.com/jcfug8/ai-writer/protos"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	log "github.com/sirupsen/logrus"
)

// Opts - Options for control service
type Opts struct {
	AIAddr      string
	PersistAddr string
	Addr        string
	AssetsDir   string
}

// Error -
type Error struct {
	message string
}

// Service - Control Service
type Service struct {
	opts          *Opts
	router        *mux.Router
	fileServer    http.Handler
	httpServer    *http.Server
	persistClient pb.PersistClient
	sessionStore  sessions.Store
}

// NewService - Takes Opts and returns an initialized control service
func NewService(p pb.PersistClient, opts *Opts) *Service {
	s := &Service{
		opts:          opts,
		router:        mux.NewRouter().StrictSlash(true),
		fileServer:    http.FileServer(http.Dir(opts.AssetsDir)),
		persistClient: p,
		sessionStore:  sessions.NewCookieStore([]byte("TEST_KEY")),
	}

	// set up api routes
	s.router.Methods("POST").Path("/api/user").HandlerFunc(s.createUser)
	s.router.Methods("GET").Path("/api/user").HandlerFunc(s.getUser)
	s.router.Methods("POST").Path("/api/auth").HandlerFunc(s.authUser)
	s.router.Methods("GET").Path("/api/book").HandlerFunc(s.getBook)
	s.router.Methods("POST").Path("/api/book").HandlerFunc(s.createBook)

	// set up static file routes
	s.router.Methods("GET").PathPrefix("/assets/").Handler(s.fileServer)
	s.router.Methods("GET").Path("/").HandlerFunc(s.serveIndex)

	// options
	s.router.Methods("OPTIONS").HandlerFunc(s.options)

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

func (s *Service) options(w http.ResponseWriter, r *http.Request) {
	log.Info("Options Hit")
}

func (s *Service) serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(s.opts.AssetsDir, "index.html"))
	return
}

func (s *Service) createUser(w http.ResponseWriter, r *http.Request) {
	req := &pb.CreateUserRequest{}

	// get request data
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		e, _ := json.Marshal(err)
		w.Write(e)
		return
	}

	// validate data
	validationErrs := []string{}
	if req.GetEmail() == "" {
		validationErrs = append(validationErrs, "invalid email")
	}

	if req.GetPassword() == "" {
		validationErrs = append(validationErrs, "invalid password")
	}

	if len(validationErrs) != 0 {
		errors.Write(w, http.StatusBadRequest, validationErrs...)
		return
	}

	// send request
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	res, err := s.persistClient.CreateUser(ctx, req)
	cancel()
	if err != nil {
		errors.WriteFromError(w, err)
		return
	}

	replies.Write(w, http.StatusCreated, res)
}

func (s *Service) authUser(w http.ResponseWriter, r *http.Request) {
	req := &pb.GetUserAuthenticateRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		e, _ := json.Marshal(err)
		w.Write(e)
		return
	}

	// validate data
	validationErrs := []string{}
	if req.GetEmail() == "" {
		validationErrs = append(validationErrs, "invalid email")
	}

	if req.GetPassword() == "" {
		validationErrs = append(validationErrs, "invalid password")
	}

	if len(validationErrs) != 0 {
		errors.Write(w, http.StatusBadRequest, validationErrs...)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	res, err := s.persistClient.GetUserAuthenticate(ctx, req)
	cancel()
	if err != nil {
		errors.WriteFromError(w, err)
		return
	}

	session, _ := s.sessionStore.Get(r, "test-session-name")
	session.Values["id"] = res.GetId()
	// need to check error s
	session.Save(r, w)

	replies.Write(w, http.StatusCreated, res)
}

func (s *Service) getUser(w http.ResponseWriter, r *http.Request) {
	session, _ := s.sessionStore.Get(r, "test-session-name")
	log.Info(session.Values)

	req := &pb.GetUserRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		e, _ := json.Marshal(err)
		w.Write(e)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	res, err := s.persistClient.GetUser(ctx, req)
	cancel()
	if err != nil {
		errors.WriteFromError(w, err)
		return
	}

	replies.Write(w, http.StatusOK, res)
}

func (s *Service) getBook(w http.ResponseWriter, r *http.Request) {
	req := &pb.GetBookRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		e, _ := json.Marshal(err)
		w.Write(e)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	res, err := s.persistClient.GetBook(ctx, req)
	cancel()
	if err != nil {
		errors.WriteFromError(w, err)
		return
	}

	replies.Write(w, http.StatusOK, res)
}

func (s *Service) createBook(w http.ResponseWriter, r *http.Request) {
	req := &pb.CreateBookRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		e, _ := json.Marshal(err)
		w.Write(e)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	res, err := s.persistClient.CreateBook(ctx, req)
	cancel()
	if err != nil {
		errors.WriteFromError(w, err)
		return
	}

	replies.Write(w, http.StatusCreated, res)
}
