package main

import (
	"context"
	"github.com/Smolvika/notebook.git"
	"github.com/Smolvika/notebook.git/pkg/handler"
	"github.com/Smolvika/notebook.git/pkg/repository"
	service1 "github.com/Smolvika/notebook.git/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("errors loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: os.Getenv("db_password"),
	})
	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
	}
	repos := repository.New(db)
	service := service1.New(repos)
	handlers := handler.New(service)
	srv := new(notebook.Server)
	go func() {
		if err := srv.Run("8080", handlers.InitRouters()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("notebookApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("notebookApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("error occured on db connection close: %s", err.Error())
	}
}
