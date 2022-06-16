package postgres

import (
	"blog-app/model/users"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os/user"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (a AuthPostgres) CreateUser(u users.User) {
	query := fmt.Sprintf(`
INSERT INTO %s (
	name,
	surname,
	father_name,
	nickname,
	password_hash
) VALUES ($1,$2,$3,$4,$5);`, usersTable)

	_, err := a.db.Query(query, u.Name, u.Surname, u.FatherName, u.Nickname, u.PasswordHash)
	if err != nil {
		log.Println(err)
	}
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
