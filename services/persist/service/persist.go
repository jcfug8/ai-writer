package service

import (
	"context"
	"database/sql"
	"net"

	pb "github.com/jcfug8/ai-writer/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	_ "github.com/go-sql-driver/mysql"
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

	s.listener, err = net.Listen("tcp", s.opts.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.server = grpc.NewServer()
	pb.RegisterPersistServer(s.server, s)

	log.Infof("perist is listening at %s", s.opts.Addr)
	return s.server.Serve(s.listener)
}

func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	log.Info("Geting User")
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Service) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookReply, error) {
	log.Info("Geting Book")
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Service) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookReply, error) {
	log.Info("Creating Book")
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
