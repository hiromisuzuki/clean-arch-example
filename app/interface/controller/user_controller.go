package controller

import (
	"net/http"
	"strconv"

	"github.com/hiromisuzuki/clean-arch-example/app/interface/database"
	"github.com/hiromisuzuki/clean-arch-example/app/model"
	"github.com/hiromisuzuki/clean-arch-example/app/usecase"
)

//usecaseのInteractorを操作するところ

//UserController UserController
type UserController struct {
	Interactor usecase.UserInteractor
	Context    Context
}

//NewUserController NewUserController
func NewUserController(sqlHandler database.SQLHandler, context Context) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
		Context: context,
	}
}

//ここでのメソッドはURIと対になる

// Create godoc
// @Summary [sample] Create a User
// @Description [sample] Create a User
// @Accept json
// @Produce  json
// @Param mail_address body email true "User Mail Address"
// @Success 200 {object} model.User
// @Failure 400 {object} infrastructure.Error
// @Failure 500 {object} infrastructure.Error
// @Router /users [post]
func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	u := model.User{}
	if err := c.Context.Bind(r, &u); err != nil {
		c.Context.StatusBadRequest(w, err)
	} else if _, err := c.Interactor.Add(u); err != nil {
		c.Context.StatusInternalServerError(w, err)
		return
	}
	c.Context.StatusCreated(w)
}

// List godoc
// @Summary [sample] Show Users
// @Description [sample] get Users
// @Produce  json
// @Param page query number false "page of results to return. default is 1."
// @Param per_page query number false "number of records to return in one request, specified as an integer from 1 to 100. default is 100."
// @Success 200 {object} model.Users
// @Failure 500 {object} infrastructure.Error
// @Router /users [get]
func (c *UserController) List(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	pagination := NewPagination(q.Get(`page`), q.Get(`per_page`))

	u, err := c.Interactor.Users(pagination)
	if err != nil {
		c.Context.StatusInternalServerError(w, err)
		return
	}
	c.Context.StatusOK(w, u)
}

// Get godoc
// @Summary [sample] Show a User
// @Description [sample] get User by ID
// @Accept  json
// @Produce  json
// @Param  id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} infrastructure.Error
// @Failure 500 {object} infrastructure.Error
// @Router /users/{id} [get]
func (c *UserController) Get(w http.ResponseWriter, r *http.Request) {
	var userID int
	param, err := c.Context.URLParam(r, "userID")
	if err == nil {
		userID, err = strconv.Atoi(param)
	}
	if err != nil {
		c.Context.StatusBadRequest(w, err)
		return
	}
	var u model.User
	u, err = c.Interactor.UserByID(userID)
	if err != nil {
		c.Context.StatusInternalServerError(w, err)
		return
	}
	c.Context.StatusOK(w, u)
}
