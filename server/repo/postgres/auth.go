package postgres

import (
	"blog-app/model/users"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (a AuthPostgres) CreateUser(u users.User) (id int) {
	query := fmt.Sprintf(`
INSERT INTO %s (
	name,
	surname,
	father_name,
	nickname,
	password_hash
) VALUES ($1,$2,$3,$4,$5)
RETURNING id;`, usersTable)

	row := a.db.QueryRow(query, u.Name, u.Surname, u.FatherName, u.Nickname, u.PasswordHash)
	err := row.Scan(&id)
	if err != nil {
		log.Println(err)
		return 0
	}

	return
}

func (a AuthPostgres) GetUser(username, pw string) (u users.User) {
	query := fmt.Sprintf(`
SELECT id, name, surname, father_name, nickname, password_hash
	FROM %s
	WHERE nickname=$1 AND password_hash=$2;`, usersTable)

	row := a.db.QueryRow(query, username, pw)
	err := row.Scan(&u.Id, &u.Name, &u.Surname, &u.FatherName, &u.Nickname, &u.PasswordHash)
	if err != nil {
		log.Println(err)
		return users.User{}
	}

	return
}

func (a AuthPostgres) GetUserById(id int) (u users.User) {
	query := fmt.Sprintf(`
SELECT id, name, surname, father_name, nickname, password_hash
	FROM %s
	WHERE id=$1;`, usersTable)

	row := a.db.QueryRow(query, id)
	err := row.Scan(&u.Id, &u.Name, &u.Surname, &u.FatherName, &u.Nickname, &u.PasswordHash)
	if err != nil {
		log.Println(err)
		return users.User{}
	}

	return
}
