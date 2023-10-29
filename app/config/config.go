package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	APP_PORT    int
	SECRET_KEY string
}

func LoadAPP() *AppConfig {
	var result = new(AppConfig)

	godotenv.Load("ENV")

	// var isRead = false
	if v, found := os.LookupEnv("APP_PORT"); found {
		port, err := strconv.Atoi(v)
		if err != nil {
			logrus.Error("Config: invalid port value,", err.Error())
			return nil
		}
		// isRead = true
		result.APP_PORT = port
	}
	return result
}