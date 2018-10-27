package interfaces

import "user_tdd/user/model"

type UserRepository interface {
	List() ([]*model.User, error)
	Store(user *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindById(id string) (*model.User, error)
	Delete(id string) error
	//Search(query string) ([]*model.User, error)
}
