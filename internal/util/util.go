package util

import (
	"fmt"
	"os"
	"strconv"
)

func StringToUint16(value string) uint16 {
	uInt64, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		panic(fmt.Errorf("failed to convert string(%s) to uint16 with: %+v", value, err))
	}
	return uint16(uInt64)
}

func GetEnvBool(value string, fallback bool) bool {
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
