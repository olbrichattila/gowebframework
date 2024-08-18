package env

import (
	"os"

	"github.com/joho/godotenv"
)

type Enver interface {
	Construct()
	Get(string) string
}

type Env struct {
}

func New() Enver {
	return &Env{}
}

const (
	envFileName = ".env.migrator"
)

func (*Env) Construct() {
	_, err := os.Stat(envFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}

		panic(err)
	}

	if err := godotenv.Load(envFileName); err != nil {
		panic(err)
	}
}

func (*Env) Get(name string) string {
	return os.Getenv(name)
}
