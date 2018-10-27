package service

import (
	"errors"
	"user_tdd/common"
	"user_tdd/user/interfaces"
	"user_tdd/user/model"
)

type UserService struct {
	repo interfaces.UserRepository
}

//NewUserService create new service
func NewUserService(repo interfaces.UserRepository) interfaces.UserService {
	return &UserService{
		repo: repo,
	}
}

//check if user is exist in database
func (self *UserService) IsDuplicatedEmail(email string) bool {
	user, _ := self.repo.FindByEmail(email)
	if user != nil {
		return true
	}

	return false
}

//Register a user
func (self *UserService) Register(user *model.User) (*model.User, error) {
	//@todo check if dublicated email or not
	return self.repo.Store(user)
}

//Delete a user
func (self *UserService) Delete(id string) error {
	user, err := self.repo.FindById(id)
	if err != nil {
		return err
	}
	if user.IsAdmin {
		return errors.New(common.ADMIN_USER_CAN_NOT_BE_DELETED)
	}
	return self.repo.Delete(id)
}
