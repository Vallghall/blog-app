package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	usersTable = "users"
	postsTable = "posts"
)

type Configs struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

func New(c *Configs) *sqlx.DB {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%s",
			c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode))
	log.Println(c.Host)
	if err != nil {
		log.Fatalln(err)
	}

	initialize(db)
	return db
}

func NewFromURL(dbURL string) *sqlx.DB {
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
