package repository

import (
	"errors"
	"strings"
	"user_tdd/common"
	"user_tdd/user/interfaces"
	"user_tdd/user/model"
)

type InMemoryUsersRepository struct {
	usersMap map[string]*model.User
}

func NewInMemoryRepository() interfaces.UserRepository {
	return &InMemoryUsersRepository{setFakeUsers()}
}

// Fake repository data
func setFakeUsers() map[string]*model.User {
	m := make(map[string]*model.User)
	m["1"] = &model.User{Id: "1", Name: "ALi", Email: "test1@test.com", Describtion: "describtion1", Age: 20, IsAdmin: true}
	m["2"] = &model.User{Id: "2", Name: "Ahmed", Email: "test2@test.com", Describtion: "describtion2", Age: 20, IsAdmin: true}
	m["3"] = &model.User{Id: "3", Name: "Mohamed", Email: "test3@test.com", Describtion: "describtion3", Age: 20, IsAdmin: true}
	return m

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
