package config

import (
	"fmt"
	"os"
	"strconv"

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
			Port:     stringToUint16(envConfig["TRANSMISSION_PORT"]),
			Protocol: envConfig["TRANSMISSION_PROTOCOL"],
			Username: envConfig["TRANSMISSION_USERNAME"],
			Password: envConfig["TRANSMISSION_PASSWORD"],
		},
		RestServer: RestServer{
			Port:        stringToUint16(envConfig["SERVER_PORT"]),
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
	isProduction := getEnvBool("IS_PROD", false)
	if isProduction {
		configSource = "config/PROD.env"
	} else {
		configSource = "config/LOCAL.env"
	}
	return configSource
}

func stringToUint16(value string) uint16 {
	uInt64, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		panic(fmt.Errorf("failed to convert string(%s) to uint16 with: %+v", value, err))
	}
	return uint16(uInt64)
}

func getEnvBool(value string, fallback bool) bool {
	envValue := os.Getenv(value)
	if len(envValue) == 0 {
		return fallback
	}
	valueBool, err := strconv.ParseBool(envValue)
	if err != nil {
		fmt.Printf("failed to parse string(%s) to bool with: %+v", value, err)
		return fallback
	}
	return valueBool
}
