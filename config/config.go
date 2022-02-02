package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	appPort string
}

var configuration config

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No env file added default port is 5006")
	}
	port := os.Getenv("PORT")
	configuration.appPort = port
}

func GetAppPort() string {
	if configuration.appPort == "" {
		return "5006"
	}
	_, err := strconv.Atoi(configuration.appPort)
	if err != nil {
		return "5006"
	}
	return configuration.appPort
}
