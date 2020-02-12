package service

import (
	"context"
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jcfug8/ai-writer/commons/errors"
	"github.com/jcfug8/ai-writer/commons/replies"
	pb "github.com/jcfug8/ai-writer/protos"
	log "github.com/sirupsen/logrus"
)

func (s *Service) serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(s.opts.AssetsDir, "index.html"))
	return
}

func (s *Service) options(w http.ResponseWriter, r *http.Request) {
	log.Info("Options Hit")
	w.Write([]byte("{}"))
}

func (s *Service) createAuthenticatedSession(w http.ResponseWriter, r *http.Request) {
	log.Info("logging in")
	req := &pb.GetUserAuthenticateRequest{}
	res := &pb.UserData{}

	// authenticate
	if userData, err := s.getUserData(w, r); userData != nil || err != nil {
		replies.Write(w, http.StatusCreated, userData)
		return
	}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		e, _ := json.Marshal(err)
		w.Write(e)
		return
	}

	// validate data
	if ok := errors.Validate(w, func(errs *errors.ValidationErrors) {
		if req.GetEmail() == "" {
			errs.Add("invalid email")
		}

		if req.GetPassword() == "" {
			errs.Add("invalid password")
		}
	}); !ok {
		return
	}

	forwardUnaryRequest(w, r, func(ctx context.Context) error {
		res, err = s.persistClient.GetUserAuthenticate(ctx, req)
		return err
	})
	if err != nil {
		log.Errorf("persist get user authentice request failed: %s", err)
		return
	}

	session, _ := s.sessionStore.Get(r, sessionName)
	session.Values[userDataKey] = res
	session.Save(r, w)

	replies.Write(w, http.StatusCreated, res)
}

func (s *Service) deleteAuthenticatedSession(w http.ResponseWriter, r *http.Request) {
	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		log.Errorf("errored while getting session data: %s", err)
		errors.Write(w, http.StatusInternalServerError, "error while deauthenticating session")
		return
	}

	delete(session.Values, userDataKey)
	s.sessionStore.Save(r, w, session)
}

func (s *Service) isLoggedIn(w http.ResponseWriter, r *http.Request) {
	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		log.Errorf("errored while getting session data to check if logged in: %s", err)
		errors.Write(w, http.StatusInternalServerError, "error checking if logged in")
		return
	}

	data, _ := session.Values[userDataKey]
	replies.Write(w, http.StatusOK, data)
}

func (s *Service) createUser(w http.ResponseWriter, r *http.Request) {
	var err error
	req := &pb.CreateUserRequest{}
	res := &pb.CreateUserReply{}

	// get request data
	if err = getRequestData(w, r, req); err != nil {
		log.Errorf("unable to parse create user request: %s", err)
		return
	}

	// validate data
	if ok := errors.Validate(w, func(errs *errors.ValidationErrors) {
		if req.GetEmail() == "" {
			errs.Add("invalid email")
		}

		if req.GetPassword() == "" {
			errs.Add("invalid password")
		}

		if req.GetFirstname() == "" {
			errs.Add("firstname cannot be empty")
		}

		if req.GetLastname() == "" {
			errs.Add("lastname cannot be empty")
		}
	}); !ok {
		return
	}

	// send request
	forwardUnaryRequest(w, r, func(ctx context.Context) error {
		res, err = s.persistClient.CreateUser(ctx, req)
		return err
	})
	if err != nil {
		log.Errorf("persist create user request failed: %s", err)
		return
	}

	replies.Write(w, http.StatusCreated, res)
}

func (s *Service) getUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var userData *pb.UserData
	req := &pb.GetUserRequest{}
	res := &pb.UserData{}

	// authenticate
	if userData = s.authenticate(w, r); userData == nil {
		return
	}

	req.Id = userData.GetId()

	forwardUnaryRequest(w, r, func(ctx context.Context) error {
		res, err = s.persistClient.GetUser(ctx, req)
		return err
	})
	if err != nil {
		log.Errorf("persist get user request failed: %s", err)
		return
	}

	replies.Write(w, http.StatusOK, res)
}

func (s *Service) getBook(w http.ResponseWriter, r *http.Request) {
	var err error
	var userData *pb.UserData
	req := &pb.GetBookRequest{}
	res := &pb.GetBookReply{}

	// authenticate
	if userData = s.authenticate(w, r); userData == nil {
		return
	}

	vars := mux.Vars(r)
	stringID := vars["id"]

	// validate data
	if ok := errors.Validate(w, func(errs *errors.ValidationErrors) {
		req.Id, err = strconv.ParseInt(stringID, 10, 64)
		if err != nil {
			errs.Add("book id must be a number")
		}
	}); !ok {
		return
	}

	req.UserId = userData.GetId()

	forwardUnaryRequest(w, r, func(ctx context.Context) error {
		res, err = s.persistClient.GetBook(ctx, req)
		return err
	})
	if err != nil {
		log.Errorf("persist get book request failed: %s", err)
		return
	}

	replies.Write(w, http.StatusOK, res)
}

func (s *Service) getBooks(w http.ResponseWriter, r *http.Request) {
	var err error
	var userData *pb.UserData
	req := &pb.ListBooksRequest{}
	res := &pb.ListBooksReply{}

	// authenticate
	if userData = s.authenticate(w, r); userData == nil {
		return
	}

	req.UserId = userData.GetId()

	forwardUnaryRequest(w, r, func(ctx context.Context) error {
		res, err = s.persistClient.ListBooks(ctx, req)
		return err
	})
	if err != nil {
		log.Errorf("persist get book request failed: %s", err)
		return
	}

	replies.Write(w, http.StatusOK, res)
}

func (s *Service) createBook(w http.ResponseWriter, r *http.Request) {
	var err error
	var userData *pb.UserData
	req := &pb.CreateBookRequest{}
	res := &pb.CreateBookReply{}

	// authenticate
	if userData = s.authenticate(w, r); userData == nil {
		return
	}

	req.UserId = userData.GetId()

	forwardUnaryRequest(w, r, func(ctx context.Context) error {
		res, err = s.persistClient.CreateBook(ctx, req)
		return err
	})
	if err != nil {
		log.Errorf("persist get book request failed: %s", err)
		return
	}

	replies.Write(w, http.StatusCreated, res)
}

func (s *Service) updateBook(w http.ResponseWriter, r *http.Request) {
	var err error
	var userData *pb.UserData
	req := &pb.UpdateBookRequest{}
	res := &pb.UpdateBookReply{}

	// authenticate
	if userData = s.authenticate(w, r); userData == nil {
		return
	}

	req.UserId = userData.GetId()

	// get request data
	if err = getRequestData(w, r, req); err != nil {
		log.Errorf("unable to parse update book request: %s", err)
		return
	}

	forwardUnaryRequest(w, r, func(ctx context.Context) error {
		res, err = s.persistClient.UpdateBook(ctx, req)
		return err
	})
	if err != nil {
		log.Errorf("persist update book request failed: %s", err)
		return
	}

	replies.Write(w, http.StatusOK, res)
}

func (s *Service) deleteBook(w http.ResponseWriter, r *http.Request) {
	var err error
	var userData *pb.UserData
	req := &pb.DeleteBookRequest{}
	res := &pb.DeleteBookReply{}

	// authenticate
	if userData = s.authenticate(w, r); userData == nil {
		return
	}

	req.UserId = userData.GetId()

	// get request data
	if err = getRequestData(w, r, req); err != nil {
		log.Errorf("unable to parse delete book request: %s", err)
		return
	}

	forwardUnaryRequest(w, r, func(ctx context.Context) error {
		res, err = s.persistClient.DeleteBook(ctx, req)
		return err
	})
	if err != nil {
		log.Errorf("persist delete book request failed: %s", err)
		return
	}

	replies.Write(w, http.StatusOK, res)
}
