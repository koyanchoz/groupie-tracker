package delivery

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotFound         = errors.New("page doesn't exist")
	ErrBadRequest       = errors.New("bad request")
	ErrEmptyInput       = errors.New("no input")
	ErrMethodNotAllowed = errors.New("method not allowed")
	ErrServer           = errors.New("internal server error")
)

func (h *Handler) ErrorHandler(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	errorText := fmt.Sprintf("%d %s\n%v", status, http.StatusText(status), err)
	h.ErrT.Execute(w, errorText)
}
