package main

import (
	"forum/internal/server"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func main() {

	godotenv.Load("cmd/app/.env")

	port := os.Getenv("PORT")
	// fmt.Println(fmt.Sprintf("%s:%s", "0.0.0.0", port))
	if port == "" {
		port = "3001"
	}
	log.Println(port, "PORT skr")
	config := server.NewConfig(&port)
	server := server.New(config) // creating new instance based on the 'config'
	// Starting the Server
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
