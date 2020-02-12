package errors

import (
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/grpc/status"
)

type ValidationErrors struct {
	errs []string
}

func (v *ValidationErrors) Add(err string) {
	v.errs = append(v.errs, err)
}

func (v *ValidationErrors) HasErrors() bool {
	return len(v.errs) != 0
}

func (v *ValidationErrors) GetErrors() []string {
	return v.errs
}

// Error -
type Error struct {
	StatusText string   `json:"status_text"`
	Status     int      `json:"status"`
	Messages   []string `json:"messages"`
}

// Bytes - returns the json error in string
func (e Error) Error() string {
	return string(e.Bytes())
}

// Bytes - returns the json error in byte array
func (e Error) Bytes() []byte {
	jsonB, err := json.Marshal(e)
	if err != nil {
		log.Panicf("Unable to marshal error: %s", err)
	}
	return jsonB
}

// New - return an error with a correctly encoded json string inside
func New(httpStatus int, messages ...string) Error {
	return Error{
		Status:     httpStatus,
		StatusText: http.StatusText(httpStatus),
		Messages:   messages,
	}
}

// WriteFromError - Sets the header and writes the error message from and error interface
func WriteFromError(w http.ResponseWriter, grpcError error) {
	st, ok := status.FromError(grpcError)
	if !ok {
		log.Panicf("Unable to parse error to grpc status - %s", grpcError)
	}

	e := Error{}
	err := json.Unmarshal([]byte(st.Message()), &e)
	if err != nil {
		log.Panicf("Unable to marshal error - %s: %s", st.Message(), err)
	}
	w.WriteHeader(e.Status)
	w.Write([]byte(st.Message()))
}

// Write - Sets the header and writes the error message
func Write(w http.ResponseWriter, httpStatus int, messages ...string) {
	w.WriteHeader(httpStatus)
	w.Write(New(httpStatus, messages...).Bytes())
}

// Write - uses a user defined function to validate
func Validate(w http.ResponseWriter, f func(errs *ValidationErrors)) bool {
	errs := &ValidationErrors{
		errs: []string{},
	}

	f(errs)
	if !errs.HasErrors() {
		return true
	}

	Write(w, http.StatusBadRequest, errs.GetErrors()...)
	return false
}
