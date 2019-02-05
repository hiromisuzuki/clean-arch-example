package controller

import (
	"errors"
	"reflect"
	"testing"

	"github.com/hiromisuzuki/clean-arch-example/app/model"
)

//see:http://haya14busa.com/golang-how-to-write-mock-of-interface-for-testing/

//fakeUserController fakeUserController
type fakeUserController struct {
	Interactor fakeUserInteractor
}

//fakeUserInteractor fakeUserInteractor
type fakeUserInteractor struct {
	Add      func(model.User) (int, error)
	Users    func() (model.Users, error)
	UserByID func(int) (model.User, error)
}

func TestCreateSuccess(t *testing.T) {
	lastInsertID := 1
	c := &fakeUserController{
		Interactor: fakeUserInteractor{
			Add: func(model.User) (int, error) {
				return lastInsertID, nil
			},
		},
	}
	user := model.User{MailAddress: `test@clean-arch-example.com`}
	expect := lastInsertID
	actual, err := c.Interactor.Add(user)
	if err != nil {
		t.Error(`an exception occurred`)
	}
	if expect != actual {
		t.Error(`expected value is different from actual value`)
	}
}
func TestCreateFailure(t *testing.T) {
	err := errors.New(`it was possible to set an invalid value`)
	c := &fakeUserController{
		Interactor: fakeUserInteractor{
			Add: func(model.User) (int, error) {
				return 0, err
			},
		},
	}
	user := model.User{MailAddress: `***invalid!***`}
	expect := err
	_, actual := c.Interactor.Add(user)
	if expect != actual {
		t.Error(`it was possible to set an invalid value`)
	}
}
func TestGetSuccess(t *testing.T) {
	userID := 1
	c := &fakeUserController{
		Interactor: fakeUserInteractor{
			UserByID: func(id int) (model.User, error) {
				user := model.User{ID: userID, MailAddress: `test@clean-arch-example.com`}
				return user, nil
			},
		},
	}
	expect := userID
	user, err := c.Interactor.UserByID(userID)
	if err != nil {
		t.Error(`an exception occurred`)
	}
	actual := user.ID
	if expect != actual {
		t.Error(`expected value is different from actual value`)
	}
}
func TestListSuccess(t *testing.T) {
	users := model.Users{}
	users.User = append(users.User, &model.User{ID: 1, MailAddress: `test1@clean-arch-example.com`})
	users.User = append(users.User, &model.User{ID: 2, MailAddress: `test2@clean-arch-example.com`})
	users.User = append(users.User, &model.User{ID: 3, MailAddress: `test3@clean-arch-example.com`})
	c := &fakeUserController{
		Interactor: fakeUserInteractor{
			Users: func() (model.Users, error) {
				return users, nil
			},
		},
	}
	expect := users
	actual, err := c.Interactor.Users()
	if err != nil {
		t.Error(`an exception occurred`)
	}
	if !reflect.DeepEqual(expect, actual) {
		t.Error(`expected value is different from actual value`)
	}
}
