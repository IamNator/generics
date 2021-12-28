package storage

import "generics/model"

type UserDatabase interface {
	CreateUser(u model.User) (*model.User, error)
	GetUser(id int) (*model.User, error)
	RemoveUser(id int) error
}

type User struct {
	s *Storage
}

func NewUser(s *Storage) *User {
	return &User{
		s: s,
	}
}

func (u *User) CreateUser(user model.User) (*model.User, error) {
	useR, er := Create(u.s, user)
	if er != nil {
		return nil, er
	}

	return useR, nil
}

func (u *User) GetUser(id int) (*model.User, error) {
	var user model.User
	user_, er := Get(u.s, user, id)
	if er != nil {
		return nil, er
	}

	return user_, nil
}

func (u *User) RemoveUser(id int) error {
	var user model.User
	er := Delete(u.s, user, id)
	if er != nil {
		return er
	}
	return nil
}
