package main

import (
	"blog-app/handlers"
	"blog-app/repo"
	"blog-app/repo/postgres"
	"blog-app/service"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	db := postgres.New(os.Getenv("DB_URL"))
	r := repo.New(db)
	s := service.New(r)
	_ = handlers.New(s)

}
