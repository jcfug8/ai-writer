package service

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jcfug8/ai-writer/commons/errors"
	pb "github.com/jcfug8/ai-writer/protos"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	log "github.com/sirupsen/logrus"
)

func init() {
	gob.Register(&pb.UserData{})
}

const (
	sessionName = "ai_writer"
	userDataKey = "userData"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, userData *pb.UserData)

type unaryReqCallback func(context.Context) error

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
	aiClient      pb.AIClient
	sessionStore  sessions.Store
}

// NewService - Takes Opts and returns an initialized control service
func NewService(p pb.PersistClient, a pb.AIClient, opts *Opts) *Service {
	s := &Service{
		opts:          opts,
		router:        mux.NewRouter().StrictSlash(true),
		fileServer:    http.FileServer(http.Dir(opts.AssetsDir)),
		persistClient: p,
		aiClient:      a,
		sessionStore:  sessions.NewCookieStore([]byte("TEST_KEY")),
	}

	s.router.Use(s.genaralMiddleware)

	// set up api routes
	s.router.Methods("POST").Path("/api/session").HandlerFunc(s.createAuthenticatedSession)
	s.router.Methods("DELETE").Path("/api/session").HandlerFunc(s.deleteAuthenticatedSession)
	s.router.Methods("GET").Path("/api/session").HandlerFunc(s.isLoggedIn)
	s.router.Methods("POST").Path("/api/user").HandlerFunc(s.createUser)
	s.router.Methods("GET").Path("/api/user").HandlerFunc(s.getUser)
	s.router.Methods("GET").Path("/api/book/{id}").HandlerFunc(s.getBook)
	s.router.Methods("GET").Path("/api/books").HandlerFunc(s.getBooks)
	s.router.Methods("DELETE").Path("/api/book").HandlerFunc(s.deleteBook)
	s.router.Methods("PUT").Path("/api/book").HandlerFunc(s.updateBook)
	s.router.Methods("POST").Path("/api/book").HandlerFunc(s.createBook)
	s.router.Methods("POST").Path("/api/getsimple").HandlerFunc(s.getSimple)

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

func (s *Service) genaralMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}

func getRequestData(w http.ResponseWriter, r *http.Request, req interface{}) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		errors.Write(w, http.StatusBadRequest, "unable to parse request")
		return err
	}
	return nil
}

func forwardUnaryRequest(w http.ResponseWriter, r *http.Request, f unaryReqCallback) error {
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	err := f(ctx)
	cancel()
	if err != nil {
		errors.WriteFromError(w, err)
		return err
	}
	return nil
}

func (s *Service) authenticate(w http.ResponseWriter, r *http.Request) *pb.UserData {
	userData, err := s.getUserData(w, r)
	if err != nil {
		return nil
	}

	if userData == nil {
		errors.Write(w, http.StatusUnauthorized, "your session is not authenticated")
		return nil
	}

	return userData
}

func (s *Service) getUserData(w http.ResponseWriter, r *http.Request) (*pb.UserData, error) {
	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		log.Errorf("errored while getting session data: %s", err)
		errors.Write(w, http.StatusInternalServerError, "error while authenticating session")
		return nil, err
	}

	data, _ := session.Values[userDataKey]
	if data == nil {
		return nil, nil
	}

	userData, ok := data.(*pb.UserData)
	if !ok {
		errors.Write(w, http.StatusInternalServerError, "unable to parse session data")
		return nil, err
	}
	return userData, nil
}
