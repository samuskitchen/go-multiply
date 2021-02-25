package main

import (
	"go-multiply/infrastructure/config"
	"log"
	"os"
)

func main() {
	log.Println("stating API cmd")
	port := os.Getenv("API_PORT")
	config.Start(port)
}