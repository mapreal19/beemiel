package envs

import (
	"os"
)

var envKey string

func Init(key string) {
	if key == "" {
		envKey = "BEEGO_ENV"
	} else {
		envKey = key
	}
}

func IsProduction() bool {
	return os.Getenv(envKey) == "production"
}

func IsDevelopment() bool {
	return os.Getenv(envKey) == "development"
}

func IsTest() bool {
	return os.Getenv(envKey) == "test"
}

func IsCI() bool {
	return os.Getenv(envKey) == "ci"
}
