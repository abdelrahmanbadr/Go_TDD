package repository

import (
	"errors"
	"strings"
	"user_tdd/common"
	"user_tdd/users/interfaces"
	"user_tdd/users/model"
)

type InMemoryUsersRepository struct {
	usersMap map[string]*model.User
}

func NewFakeRepository() interfaces.UserRepository {
	var userMap map[string]*model.User
	return &InMemoryUsersRepository{userMap}
}

// Fake repository data
func setFakeUsers() []*model.User {
	return []*model.User{
		{Id: "1", Email: "test1@test.com", Describtion: "describtion1", Age: 20, IsAdmin: true},
		{Id: "2", Email: "test2@test.com", Describtion: "describtion2", Age: 30, IsAdmin: false},
		{Id: "3", Email: "test3@test.com", Describtion: "describtion3", Age: 40, IsAdmin: false},
		{Id: "4", Email: "test4@test.com", Describtion: "describtion4", Age: 50, IsAdmin: false},
	}

}

func (self *InMemoryUsersRepository) List() ([]*model.User, error) {
	var users []*model.User
	for _, user := range self.usersMap {
		users = append(users, user)
	}
	return users, nil
}

func (self *InMemoryUsersRepository) FindByEmail(email string) (*model.User, error) {

	for _, user := range self.usersMap {
		if strings.Contains(strings.ToLower(user.Email), email) {
			return user, nil
		}
	}
	return nil, errors.New(common.USER_NOT_FOUND)
}

//Find user by id
func (self *InMemoryUsersRepository) FindById(id string) (*model.User, error) {
	user := self.usersMap[id]
	if user == nil {
		return nil, errors.New(common.USER_NOT_FOUND)
	}
	return user, nil
}

func (self *InMemoryUsersRepository) Store(user *model.User) (*model.User, error) {

	self.usersMap[user.Id] = user
	return user, nil
}

//Delete a user
func (self *InMemoryUsersRepository) Delete(id string) error {
	_, err := self.FindById(id)
	if err != nil {
		return err
	}

	self.usersMap[id] = nil
	return nil
}

func (self *InMemoryUsersRepository) Search(query string) ([]*model.User, error) {
	var users []*model.User
	for _, user := range self.usersMap {
		if strings.Contains(strings.ToLower(user.Name), query) {
			users = append(users, user)
		}
	}
	if len(users) == 0 {
		return nil, errors.New(common.USER_NOT_FOUND)
	}

	return users, nil
}
