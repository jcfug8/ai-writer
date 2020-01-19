package service

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	pb "github.com/jcfug8/ai-writer/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	log.Infof("Creating User: req - %+v", req)

	// validate data
	validationErrs := []string{}
	if req.GetEmail() == "" {
		validationErrs = append(validationErrs, "invalid email")
	}

	if req.GetPassword() == "" {
		validationErrs = append(validationErrs, "invalid password")
	}

	if len(validationErrs) != 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("%v", validationErrs))
	}

	// begin transaction
	tx, err := s.db.Begin()
	if err != nil {
		return nil, status.Error(codes.Unavailable, fmt.Sprintf("could not create transaction: %s", err))
	}

	// check for conflicting email
	rows := tx.QueryRowContext(ctx, "SELECT id FROM users WHERE email = ?", req.GetEmail())
	var id int
	if err = rows.Scan(&id); err != sql.ErrNoRows {
		tx.Rollback()
		errMsg := "Email is already registered"
		if err != nil {
			errMsg = errMsg + ": " + err.Error()
		}
		return nil, status.Error(codes.AlreadyExists, errMsg)
	}

	// insert users
	_, err = tx.ExecContext(ctx, "INSERT INTO users(email, hashed_password) VALUES(?, ?)", req.GetEmail(), req.GetPassword())
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.Unavailable, fmt.Sprintf("could not insert user: %s", err))
	}

	// finish
	tx.Commit()
	return &pb.CreateUserReply{}, nil
}

func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	log.Infof("Getting User: req - %+v", req)
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Service) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookReply, error) {
	log.Infof("Getting Book: req - %+v", req)
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Service) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookReply, error) {
	log.Infof("Creating Book: req - %+v", req)
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
