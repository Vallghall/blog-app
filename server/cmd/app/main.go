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

	db := postgres.New(getDBConfigs())
	r := repo.New(db)
	s := service.New(r)
	g := handlers.New(s).HandleRoutes()
	log.Fatalln(g.Run(":" + os.Getenv("PORT")))
}

func getDBConfigs() *postgres.Configs {
	return &postgres.Configs{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PW"),
		SSLMode:  os.Getenv("DB_SSL"),
	}
}
