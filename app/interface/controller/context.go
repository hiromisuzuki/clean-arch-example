package controller

import "net/http"

//Context Context
type Context interface {
	URLParam(r *http.Request, key string) (string, error)
	Bind(r *http.Request, v interface{}) error

	StatusOK(w http.ResponseWriter, data interface{})
	StatusCreated(w http.ResponseWriter)
	StatusAccepted(w http.ResponseWriter)
	StatusNoContent(w http.ResponseWriter)
	StatusBadRequest(w http.ResponseWriter, err error)
	StatusInternalServerError(w http.ResponseWriter, err error)
}
