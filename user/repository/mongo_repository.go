package repository

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"user_tdd/user/interfaces"
	"user_tdd/user/model"
)

type MongoUserRepo struct {
	dbConnection *mgo.Database
}

func NewMongoUserRepo(Conn *mgo.Database) interfaces.UserRepository {

	return &MongoUserRepo{Conn}
}

func (self *MongoUserRepo) GetUserCollection() *mgo.Collection {

	return self.dbConnection.C("user")
}

func (self *MongoUserRepo) List() ([]*model.User, error) {
	var users []*model.User

	err := self.GetUserCollection().Find(nil).All(&users)
	return users, err
}

func (self *MongoUserRepo) FindByEmail(email string) (*model.User, error) {
	var newUser model.User
	err := self.GetUserCollection().Find(bson.M{"email": email}).One(&newUser)
	return &newUser, err
}

func (self *MongoUserRepo) FindById(id string) (*model.User, error) {
	var newUser model.User
	err := self.GetUserCollection().Find(bson.M{"id": id}).One(&newUser)
	return &newUser, err
}

func (self *MongoUserRepo) Store(user *model.User) (*model.User, error) {

	err := self.GetUserCollection().Insert(user)

	return user, err
}

//Delete a bookmark
func (self *MongoUserRepo) Delete(id string) error {

	return self.GetUserCollection().Remove(bson.M{"id": id})
}
