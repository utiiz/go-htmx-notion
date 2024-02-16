package utils

import (
	"os"

	"github.com/joho/godotenv"
)

var DOT_ENV_PATH = ".env"

type Env struct {
	DB_USER string
	DB_PASS string
	DB_NAME string
}

func GetEnv() (*Env, error) {
	err := godotenv.Load(DOT_ENV_PATH)
	if err != nil {
		return nil, err
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	return &Env{
		DB_USER: user,
		DB_PASS: pass,
		DB_NAME: name,
	}, nil
}
