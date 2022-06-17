package service

import (
	"blog-app/model/users"
	"blog-app/repo"
)

type Auth struct {
	r *repo.Repo
}

func NewAuthService(r *repo.Repo) *Auth {
	return &Auth{r}
}

func (a *Auth) CreateUser(user users.User) {
	a.r.CreateUser(user)
}

func (a *Auth) GetUser(username, pw string) users.User {
	//TODO implement me
	panic("implement me")
}

func (a *Auth) GetUserById(id int) users.User {
	//TODO implement me
	panic("implement me")
}
