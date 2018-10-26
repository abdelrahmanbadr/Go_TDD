package interfaces

import "user_tdd/users/model"

type UserService interface {
	IsDuplicatedEmail(email string) bool
	Register(user *model.User) (*model.User, error)
	Delete(id string) error
	//ForgotPassword(user *model.User) error
	//ChangePassword(user *model.User, password string) error
	//GetRepo() UserRepository
}
