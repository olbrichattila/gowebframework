package env

import (
	"fmt"
	"framework/internal/app/logger"
	"os"

	"github.com/joho/godotenv"
)

type Enver interface {
	Construct(l logger.Logger)
	Get(string) string
}

type Env struct {
	l      logger.Logger
	loaded bool
}

func New() Enver {
	return &Env{
		loaded: false,
	}
}

var envFileNames = []string{".env.migrator", ".env"}

func (e *Env) Construct(l logger.Logger) {
	if e.loaded {
		return
	}

	e.l = l
	for _, envFileName := range envFileNames {
		e.loadEnvIfExits(envFileName)
	}
	e.loaded = true
}

func (e *Env) loadEnvIfExits(envFileName string) {
	_, err := os.Stat(envFileName)
	if err == nil {
		if err := godotenv.Load(envFileName); err != nil {
			e.logEnvLoadError(envFileName, err.Error())
			return
		}
		return
	}

	e.logEnvLoadError(envFileName, err.Error())
}

func (e *Env) logEnvLoadError(envFileName, errorName string) {
	if e.l != nil {
		e.l.Info(
			fmt.Sprintf("Cannot load .env file %s: %s", envFileName, errorName),
		)
	}
}

func (*Env) Get(name string) string {
	return os.Getenv(name)
}
