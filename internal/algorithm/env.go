package algorithm

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	ServerPort   string
	MAPS_API_KEY string
}

var ENV *Env

func newEnv() *Env {
	port := "3333"
	if os.Getenv("SERVER_PORT") != "" {
		port = os.Getenv("SERVER_PORT")
	}

	return &Env{
		ServerPort:   port,
		MAPS_API_KEY: os.Getenv("MAPS_API_KEY"),
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	ENV = newEnv()
}
