package main

import (
	"github.com/Smolvika/notebook.git"
	"log"
)

func main() {
	srv := new(notebook.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error running http serever: %s", err.Error())
	}
}
