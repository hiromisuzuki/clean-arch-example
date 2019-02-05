package database

import (
	"reflect"
	"testing"

	"github.com/hiromisuzuki/clean-arch-example/app/model"
	"github.com/hiromisuzuki/clean-arch-example/app/usecase"
)

//see:http://haya14busa.com/golang-how-to-write-mock-of-interface-for-testing/

//fakeUserRepository fakeUserRepository
type fakeUserRepository struct {
	usecase.UserRepository //interfaceと同じメソッド型(のフェイク)をフィールド定義する
	FakeStore              func(model.User) (int, error)
	FakeFindByID           func(int) (model.User, error)
	FakeFindAll            func() (model.Users, error)
}

func (m *fakeUserRepository) Store(u model.User) (int, error) {
	return m.FakeStore(u)
}

func (m *fakeUserRepository) FindByID(id int) (model.User, error) {
	return m.FakeFindByID(id)
}

func (m *fakeUserRepository) FindAll() (model.Users, error) {
	return m.FakeFindAll()
}

func TestStore(t *testing.T) {
	lastInsertID := 1
	repo := &fakeUserRepository{
		FakeStore: func(u model.User) (int, error) {
			return lastInsertID, nil
		},
	}
	user := model.User{MailAddress: `test@clean-arch-example.com`}
	expect := lastInsertID
	actual, err := repo.Store(user)
	if err != nil {
		t.Error(`an exception occurred`)
	}
	if expect != actual {
		t.Error(`expected value is different from actual value`)
	}
}

func TestFindByID(t *testing.T) {
	userID := 1
	repo := &fakeUserRepository{
		FakeFindByID: func(id int) (model.User, error) {
			user := model.User{ID: userID, MailAddress: `test@clean-arch-example.com`}
			return user, nil
		},
	}
	user, err := repo.FindByID(userID)
	if err != nil {
		t.Error(`an exception occurred`)
	}
	expect := userID
	actual := user.ID
	if expect != actual {
		t.Error(`expected value is different from actual value`)
	}
}

func TestFindAll(t *testing.T) {
	users := model.Users{}
	users.User = append(users.User, &model.User{ID: 1, MailAddress: `test1@clean-arch-example.com`})
	users.User = append(users.User, &model.User{ID: 2, MailAddress: `test2@clean-arch-example.com`})
	users.User = append(users.User, &model.User{ID: 3, MailAddress: `test3@clean-arch-example.com`})
	repo := &fakeUserRepository{
		FakeFindAll: func() (model.Users, error) {
			return users, nil
		},
	}
	expect := users
	actual, err := repo.FindAll()
	if err != nil {
		t.Error(`an exception occurred`)
	}
	if !reflect.DeepEqual(expect, actual) {
		t.Error(`expected value is different from actual value`)
	}
}
