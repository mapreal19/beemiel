package envs

import (
	"os"
)

var envKey = "BEEGO_ENV"

func Init(key string) {
	if key == "" {
		return
	}

	envKey = key
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
