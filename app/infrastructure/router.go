package infrastructure

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/hiromisuzuki/clean-arch-example/app/interface/controller"
)

//ルーティングを定義するところ。
//外部ライブラリ(chiであったりechoであったり)を使うのでここに書く

//NewRouter NewRouter
func NewRouter(config *SQLConfig) *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get(`/`, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`clean-arch-example`))
	})

	handler := NewSQLHandler(config)
	context := NewContext()

	//API Sample: users
	userController := controller.NewUserController(handler, context)

	r.Post(`/users`, userController.Create)
	r.Get(`/users`, userController.List)
	r.Get(`/users/{userID}`, userController.Get)

	return r
}
