package usecase

import "github.com/hiromisuzuki/clean-arch-example/app/model"

//　DB寄りの抽象的なインターフェース

//UserRepository UserRepository interface
type UserRepository interface {
	Store(model.User) (int, error)
	FindByID(int) (model.User, error)
	FindAll(model.Pagination) (model.Users, error)
}
