package infrastructure

import (
	"encoding/json"
	"net/http"
)

//Response Response
type Response struct {
	Body interface{}
}

//Error Error
type Error struct {
	Message string `json:"message"`
}

//StatusOK StatusOK 200
func (r *Response) StatusOK(w http.ResponseWriter, data interface{}) {
	r.statusSuccess(w, http.StatusOK, data)
}

//StatusCreated StatusCreated 201
func (r *Response) StatusCreated(w http.ResponseWriter) {
	r.statusSuccess(w, http.StatusCreated, nil)
}

//StatusAccepted StatusAccepted 202
func (r *Response) StatusAccepted(w http.ResponseWriter) {
	r.statusSuccess(w, http.StatusAccepted, nil)
}

//StatusNoContent StatusNoContent 204
func (r *Response) StatusNoContent(w http.ResponseWriter) {
	r.statusSuccess(w, http.StatusNoContent, nil)
}

//StatusBadRequest StatusBadRequest 400
func (r *Response) StatusBadRequest(w http.ResponseWriter, err error) {
	r.statusFailure(w, http.StatusBadRequest, err)
}

//StatusInternalServerError StatusInternalServerError 500
func (r *Response) StatusInternalServerError(w http.ResponseWriter, err error) {
	r.statusFailure(w, http.StatusInternalServerError, err)
}

func (r *Response) statusSuccess(w http.ResponseWriter, status int, data interface{}) {
	r.response(w, status, data)
}

func (r *Response) statusFailure(w http.ResponseWriter, status int, err error) {
	r.response(w, status, Error{Message: err.Error()})
}

func (r *Response) response(w http.ResponseWriter, status int, body interface{}) {
	r.Body = body
	w.Header().Set(`Content-Type`, `application/json`)
	w.WriteHeader(status)
	w.Write(r.toJSON())
}

//toJSON toJSON
func (r *Response) toJSON() []byte {
	b, err := json.Marshal(r.Body)
	if err != nil {
		return ([]byte)(`{
			Message: "JSON Parse error."
		}`)
	}
	return b
}
