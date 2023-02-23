package main

import (
	"github.com/Smolvika/notebook.git"
	"github.com/Smolvika/notebook.git/pkg/handler"
	"github.com/Smolvika/notebook.git/pkg/repository"
	service2 "github.com/Smolvika/notebook.git/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("errors loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "32768",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: os.Getenv("db_password"),
	})
	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	service := service2.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(notebook.Server)
	if err := srv.Run("8080", handlers.InitRouters()); err != nil {
		log.Fatalf("error running http serever: %s", err.Error())
	}
}
