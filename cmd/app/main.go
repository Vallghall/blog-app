package main

import (
	"blog-app/handlers"
	"blog-app/repo"
	"blog-app/service"
	"github.com/jmoiron/sqlx"
)

func main() {
	r := repo.New(&sqlx.DB{})
	s := service.New(r)
	_ = handlers.New(s)

}
