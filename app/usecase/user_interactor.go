package usecase

import "github.com/hiromisuzuki/clean-arch-example/app/model"

//UserInteractor UserInteractor
type UserInteractor struct {
	UserRepository UserRepository
}

//メソッド名(Store=>Add,FindAll=>Users)が具象化している

//Add Add
func (interactor *UserInteractor) Add(u model.User) (id int, err error) {
	id, err = interactor.UserRepository.Store(u)
	return
}

//Users Users
func (interactor *UserInteractor) Users(pagination model.Pagination) (u model.Users, err error) {
	u, err = interactor.UserRepository.FindAll(pagination)
	return
}

//UserByID UserByID
func (interactor *UserInteractor) UserByID(id int) (u model.User, err error) {
	u, err = interactor.UserRepository.FindByID(id)
	return
}
