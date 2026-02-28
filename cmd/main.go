package main

import (
	"log"
	"os"

	"github.com/OlegRozh/Final-progect-todo/internal/server"
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
	dbPath := os.Getenv("TODO_DBFILE")

	store, err := storage.New(dbPath)
	if err != nil {
		log.Fatal("Failed to create storage:", err)
	}
	defer store.DB.Close()

	srv := server.New(store, port)
	srv.SetupRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("Server failed:", err)
	}

}
