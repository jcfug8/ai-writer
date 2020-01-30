package service

import (
	"context"
	"database/sql"
	"net"
	"net/http"

	"github.com/jcfug8/ai-writer/commons/errors"
	pb "github.com/jcfug8/ai-writer/protos"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

// Opts - Options for control service
type Opts struct {
	Addr string
}

// Service - Persist Service
type Service struct {
	opts     *Opts
	listener net.Listener
	server   *grpc.Server
	db       *sql.DB
}

// NewService - Takes Opts and returns an initialized control service
func NewService(opts *Opts) *Service {
	s := &Service{
		opts: opts,
	}
	return s
}

// Serve - Start up the persist service
func (s *Service) Serve() error {
	var err error
	log.Info("serving persist")

	// set up grpc server
	s.listener, err = net.Listen("tcp", s.opts.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.server = grpc.NewServer()
	pb.RegisterPersistServer(s.server, s)

	log.Infof("perist is listening at %s", s.opts.Addr)
	return s.server.Serve(s.listener)
}

func (s *Service) RegisterDatabase(db *sql.DB) {
	s.db = db
}

// CreateUser - Inserts a user record into the db. It checks and fails if the email is a duplicate.
// It also store the password as a salted hash
func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	log.Infof("Creating User: req - %+v", req)

	// begin transaction
	tx, err := s.db.Begin()
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "could not create database transaction")
	}

	// check for conflicting email
	rows := tx.QueryRowContext(ctx, "SELECT id FROM users WHERE email = ?", req.GetEmail())
	var id int
	if err = rows.Scan(&id); err != sql.ErrNoRows {
		tx.Rollback()
		log.Infof("unable to register user with email %s because the email is taken", req.GetEmail())
		return nil, errors.New(http.StatusConflict, "unable to register user")
	}

	// insert user
	hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), 10)
	if err != nil {
		tx.Rollback()
		log.Errorf("unable to hash a password: %s", err)
		return nil, errors.New(http.StatusInternalServerError, "could not hash password")
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO users(email, hashed_password) VALUES(?, ?)", req.GetEmail(), hash)
	if err != nil {
		tx.Rollback()
		log.Errorf("unable to insert user: %s", err)
		return nil, errors.New(http.StatusInternalServerError, "could not insert user")
	}

	// finish
	tx.Commit()
	return &pb.CreateUserReply{}, nil
}

// GetUser - Selects a user by the id given
func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserData, error) {
	rows := s.db.QueryRowContext(ctx, "SELECT id, email FROM users WHERE id = ?", req.GetId())
	var id int64
	var email string

	if err := rows.Scan(&id, &email); err == sql.ErrNoRows {
		log.Infof("row was not found when getting user %d: %s", req.GetId(), err)
		return nil, errors.New(http.StatusNotFound, "user not found")
	} else if err != nil {
		log.Infof("error when getting user %d: %s", req.GetId(), err)
		return nil, errors.New(http.StatusInternalServerError, "uable to retrieve user")
	}

	return &pb.UserData{
		Id:    id,
		Email: email,
	}, nil
}

// GetUserAuthenticate - Gets a users by email and password
func (s *Service) GetUserAuthenticate(ctx context.Context, req *pb.GetUserAuthenticateRequest) (*pb.UserData, error) {
	rows := s.db.QueryRowContext(ctx, "SELECT id, hashed_password FROM users WHERE email = ?", req.GetEmail())
	var id int64
	var hashedPassword string
	err := rows.Scan(&id, &hashedPassword)

	if err != nil {
		log.Infof("errored finding user with email %s: %s", req.GetEmail(), err)
		return nil, errors.New(http.StatusUnauthorized, "unable to authenticate")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.GetPassword())); err != nil {
		log.Infof("passwords did not match for user with email %s: %s", req.GetEmail(), err)
		return nil, errors.New(http.StatusUnauthorized, "not able to authenticate")
	}

	return &pb.UserData{
		Id:    id,
		Email: req.GetEmail(),
	}, nil
}

// GetBook -
func (s *Service) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookReply, error) {
	log.Infof("Getting Book: req - %+v", req)
	return nil, errors.New(http.StatusNotImplemented, "unimplemented")
}

// CreateBook -
func (s *Service) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookReply, error) {
	log.Infof("Creating Book: req - %+v", req)
	return nil, errors.New(http.StatusNotImplemented, "unimplemented")
}
