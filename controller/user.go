package controller

import "generics/model"

func (c *Controller) RegisterUser(u model.User) (model.User, error) {
	u.Password = u.Password.Hash()
	u_, er := c.user.CreateUser(u)
	if er != nil {
		return u, er
	}
	return *u_, nil
}
