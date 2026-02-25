package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/OlegRozh/Final-progect-todo/internal/api"
	"github.com/OlegRozh/Final-progect-todo/internal/storage"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Port must be set, using default 8080")
		port = "8080"
	}
	url := fmt.Sprintf(":%s", port)

	store, err := storage.New()
	if err != nil {
		log.Fatal("Failed to init storage:", err)
	}
	api.Store = store

	fileServer := http.FileServer(http.Dir("web"))
	http.Handle("/", fileServer)

	log.Printf("Listening on %s", url)
	err = http.ListenAndServe(url, nil)
	if err != nil {
		panic(err)
	}

}
