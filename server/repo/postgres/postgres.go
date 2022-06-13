package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
)

func New(dbURL string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
