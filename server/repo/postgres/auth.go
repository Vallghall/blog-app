package postgres

import (
	"blog-app/model/users"
	"github.com/jmoiron/sqlx"
	"os/user"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (a AuthPostgres) CreateUser(user users.User) {
	//TODO implement me
	panic("implement me")
}

func (a AuthPostgres) GetUser(username, pw string) users.User {
	//TODO implement me
	panic("implement me")
}

func (a AuthPostgres) GetUserById() user.User {
	//TODO implement me
	panic("implement me")
}

func (a AuthPostgres) UserExists(u user.User) bool {
	//TODO implement me
	panic("implement me")
}
