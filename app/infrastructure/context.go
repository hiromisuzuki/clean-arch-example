package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

//Context Context
type Context struct {
	Response
}

//NewContext NewContext
func NewContext() *Context {
	return new(Context)
}

//URLParam URLParam
func (c *Context) URLParam(r *http.Request, key string) (string, error) {
	if param := chi.URLParam(r, key); param != "" {
		return param, nil
	}
	return "", fmt.Errorf("invalid parameter.[%s]", key)
}

//Bind Bind
func (c *Context) Bind(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
