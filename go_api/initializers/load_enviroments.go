package initializers

import (
	"os"

	"github.com/joho/godotenv"
)

var Port, DbHost, DbPassword, DbUser, DbPort, DbName string

func LoadEnviroments() {
	godotenv.Load()
	Port = os.Getenv("PORT")
	DbHost = os.Getenv("DB_HOST")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	DbPort = os.Getenv("DB_PORT")
}
