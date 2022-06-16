package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	usersTable = "users"
	postsTable = "posts"
)

func New(dbURL string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalln(err)
	}

	initialize(db)
	return db
}

func initialize(db *sqlx.DB) {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS "users" 
(
    "id" serial PRIMARY KEY,
	"name" varchar NOT NULL,
    "surname" varchar NOT NULL,
    "father_name" varchar,
    "nickname" varchar NOT NULL UNIQUE,
    "password_hash" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "posts" 
(
    "id" serial PRIMARY KEY,
    "author_id" INTEGER NOT NULL REFERENCES users(id),
	"title" varchar NOT NULL,
    "content" varchar NOT NULL,
    "date" TIMESTAMP NOT NULL ,
    "hashtags" varchar ARRAY
)

`)
	if err != nil {
		log.Fatalln(err)
	}
}
