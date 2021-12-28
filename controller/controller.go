package controller

import (
	"generics/model"
	"generics/storage"
)

type Operation interface {
	RegisterUser(u model.User) (model.User, error)
}

type Controller struct {
	user storage.UserDatabase
}

func New(s *storage.Storage) *Controller {
	user := storage.NewUser(s)
	return &Controller{
		user: user,
	}
}
