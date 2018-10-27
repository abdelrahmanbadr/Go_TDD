package model

import "github.com/globalsign/mgo/bson"

type User struct {
	ID          bson.ObjectId `bson:"_id" json:"-"`
	Id          string        `json:"id"`
	Email       string        `json:"email" validate:"required"`
	Name        string        `json:"name" validate:"required"`
	Describtion string        `json:"describtion" validate:"required"`
	Age         int           `json:"age" validate:"required"`
	IsAdmin     bool          `json:"bool" validate:"required"`
}

func NewUser(id string, email string, name string, describtion string, age int, isAdmin bool) *User {
	return &User{
		Id:          id,
		Email:       email,
		Name:        name,
		Describtion: describtion,
		Age:         age,
		IsAdmin:     isAdmin,
	}
}
