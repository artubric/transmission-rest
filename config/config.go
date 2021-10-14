package config

import (
	"artubric/transmission-rest/internal/util"
	"fmt"

	"github.com/joho/godotenv"
)

type TransmissionServer struct {
	Host     string
	Port     uint16
	Protocol string
	Username string
	Password string
}

type RestServer struct {
	Port        uint16
	ApiBasePath string
	MainEntity  string
}

type Log struct {
	Level string
}

type Config struct {
	TransmissionServer TransmissionServer
	RestServer         RestServer
	Log                Log
}

func Load() *Config {

	envConfig, err := godotenv.Read(getConfigSource())

	if err != nil {
		panic(fmt.Errorf("failed to load .env file with: %+v", err))
	}

	config := &Config{
		TransmissionServer: TransmissionServer{
			Host:     envConfig["TRANSMISSION_HOST"],
			Port:     util.StringToUint16(envConfig["TRANSMISSION_PORT"]),
			Protocol: envConfig["TRANSMISSION_PROTOCOL"],
			Username: envConfig["TRANSMISSION_USERNAME"],
			Password: envConfig["TRANSMISSION_PASSWORD"],
		},
		RestServer: RestServer{
			Port:        util.StringToUint16(envConfig["SERVER_PORT"]),
			ApiBasePath: envConfig["SERVER_BASE_PATH"],
			MainEntity:  envConfig["SERVER_MAIN_ENTITY"],
		},
		Log: Log{
			Level: envConfig["LOG_LEVEL"],
		},
	}
	return config
}

func getConfigSource() string {
	var configSource string
	isProduction := util.GetEnvBool("IS_PROD", false)
	if isProduction {
		configSource = "config/.env.prod"
	} else {
		configSource = "config/.env.local"
	}
	return configSource
}
