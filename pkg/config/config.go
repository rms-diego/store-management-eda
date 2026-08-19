package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	PORT         string
	DATABASE_URL string
}

var Env *env

func Init() error {
	godotenv.Load()

	switch {
	case os.Getenv("PORT") == "":
		return fmt.Errorf("Environment variable 'PORT' must be declared")

	case os.Getenv("DATABASE_URL") == "":
		return fmt.Errorf("Environment variable 'DATABASE_URL' must be declared")

	default:
		Env = &env{
			PORT:         os.Getenv("PORT"),
			DATABASE_URL: fmt.Sprintf("%v/store_management", os.Getenv("DATABASE_URL")),
		}

		return nil
	}
}
