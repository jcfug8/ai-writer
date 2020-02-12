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

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO users(firstname, lastname, email, hashed_password) VALUES(?, ?, ?, ?)",
		req.GetFirstname(),
		req.GetLastname(),
		req.GetEmail(),
		hash,
	)
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
	rows := s.db.QueryRowContext(ctx, "SELECT id, email, firstname, lastname FROM users WHERE id = ?", req.GetId())
	var id int64
	var email string
	var firstname string
	var lastname string

	if err := rows.Scan(&id, &email, &firstname, &lastname); err == sql.ErrNoRows {
		log.Infof("row was not found when getting user %d: %s", req.GetId(), err)
		return nil, errors.New(http.StatusNotFound, "user not found")
	} else if err != nil {
		log.Infof("error when getting user %d: %s", req.GetId(), err)
		return nil, errors.New(http.StatusInternalServerError, "uable to retrieve user")
	}

	return &pb.UserData{
		Id:        id,
		Email:     email,
		Firstname: firstname,
		Lastname:  lastname,
	}, nil
}

// GetUserAuthenticate - Gets a users by email and password
func (s *Service) GetUserAuthenticate(ctx context.Context, req *pb.GetUserAuthenticateRequest) (*pb.UserData, error) {
	rows := s.db.QueryRowContext(ctx, "SELECT id, hashed_password, firstname, lastname FROM users WHERE email = ?", req.GetEmail())
	var id int64
	var hashedPassword string
	var firstname string
	var lastname string
	err := rows.Scan(&id, &hashedPassword, &firstname, &lastname)

	if err != nil {
		log.Infof("errored finding user with email %s: %s", req.GetEmail(), err)
		return nil, errors.New(http.StatusUnauthorized, "unable to authenticate")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.GetPassword())); err != nil {
		log.Infof("passwords did not match for user with email %s: %s", req.GetEmail(), err)
		return nil, errors.New(http.StatusUnauthorized, "not able to authenticate")
	}

	return &pb.UserData{
		Id:        id,
		Email:     req.GetEmail(),
		Firstname: firstname,
		Lastname:  lastname,
	}, nil
}

// GetBook -
func (s *Service) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookReply, error) {
	rows := s.db.QueryRowContext(ctx, "SELECT id, name, description, body FROM books WHERE id = ? AND user_id = ?", req.GetId(), req.GetUserId())
	reply := &pb.GetBookReply{}

	if err := rows.Scan(&reply.Id, &reply.Name, &reply.Description, &reply.Body); err == sql.ErrNoRows {
		log.Infof("book %d was not found for user %d: %s", req.GetId(), req.GetUserId(), err)
		return nil, errors.New(http.StatusNotFound, "book not found")
	} else if err != nil {
		log.Infof("book %d was not found for user %d: %s", req.GetId(), req.GetUserId(), err)
		return nil, errors.New(http.StatusInternalServerError, "error retrieving book")
	}

	return reply, nil
}

// ListBooks -
func (s *Service) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksReply, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, description FROM books WHERE user_id = ?", req.GetUserId())

	if err != nil {
		log.Infof("unable to query for books for user %d: %s", req.GetUserId(), err)
		return nil, errors.New(http.StatusInternalServerError, "unable get books list")
	}
	defer rows.Close()

	books := []*pb.ListBook{}

	for rows.Next() {
		book := &pb.ListBook{}
		err = rows.Scan(&book.Id, &book.Name, &book.Description)
		if err != nil {
			log.Infof("error scaning over books for user %d: %s", req.GetUserId(), err)
			return nil, errors.New(http.StatusInternalServerError, "error scanning over book list")
		}
		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		log.Infof("erros after scanning books for user %d: %s", req.GetUserId(), err)
		return nil, errors.New(http.StatusInternalServerError, "error after scanning for book list")
	}

	return &pb.ListBooksReply{
		Books: books,
	}, nil
}

// CreateBook -
func (s *Service) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookReply, error) {
	// begin transaction
	tx, err := s.db.Begin()
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "could not create database transaction")
	}

	res, err := tx.ExecContext(
		ctx,
		"INSERT INTO books(user_id, name, description, body) VALUES(?, ?, ?, ?)",
		req.GetUserId(),
		"Untitled",
		"",
		"",
	)
	if err != nil {
		tx.Rollback()
		log.Errorf("unable to create book for user %d: %s", req.GetUserId(), err)
		return nil, errors.New(http.StatusInternalServerError, "could not create book")
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Errorf("unable to retrieve id of book created for user %d: %s", req.GetUserId(), err)
		return nil, errors.New(http.StatusInternalServerError, "could not create book")
	}

	// finish
	tx.Commit()
	return &pb.CreateBookReply{
		Id: id,
	}, nil
}

// UpdateBook -
func (s *Service) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookReply, error) {
	log.Infof("update book - %+v", req)

	_, err := s.db.ExecContext(
		ctx,
		"UPDATE books SET name = ?, description = ?, body = ? WHERE id = ? AND user_id = ?",
		req.GetName(),
		req.GetDescription(),
		req.GetBody(),
		req.GetId(),
		req.GetUserId(),
	)
	if err != nil {
		log.Errorf("unable to update book %s for user %d: %s", req.GetId(), req.GetUserId(), err)
		return nil, errors.New(http.StatusInternalServerError, "could not update book")
	}

	return &pb.UpdateBookReply{}, nil
}

// DeleteBook -
func (s *Service) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookReply, error) {
	log.Infof("delete book - %+v", req)

	_, err := s.db.ExecContext(
		ctx,
		"DELETE FROM books WHERE id = ? AND user_id = ?",
		req.GetId(),
		req.GetUserId(),
	)
	if err != nil {
		log.Errorf("unable to delete book %s for user %d: %s", req.GetId(), req.GetUserId(), err)
		return nil, errors.New(http.StatusInternalServerError, "could not delete book")
	}

	return &pb.DeleteBookReply{}, nil
}
