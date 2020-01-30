package service

import (
	"context"
	"encoding/json"
	"net/http"
	"path/filepath"

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
}

func (s *Service) createAuthenticatedSession(w http.ResponseWriter, r *http.Request) {
	req := &pb.GetUserAuthenticateRequest{}
	res := &pb.UserData{}

	// authenticate
	if userData, err := s.getUserData(w, r); userData != nil || err != nil {
		return
	}

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

	replies.Write(w, http.StatusCreated, map[string]string{})
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

	// get request data
	if err = getRequestData(w, r, req); err != nil {
		log.Errorf("unable to parse get book request: %s", err)
		return
	}

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

func (s *Service) createBook(w http.ResponseWriter, r *http.Request) {
	var err error
	var userData *pb.UserData
	req := &pb.CreateBookRequest{}
	res := &pb.CreateBookReply{}

	// authenticate
	if userData = s.authenticate(w, r); userData == nil {
		return
	}

	// get request data
	if err = getRequestData(w, r, req); err != nil {
		log.Errorf("unable to parse get book request: %s", err)
		return
	}

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
