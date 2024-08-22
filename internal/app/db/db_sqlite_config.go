package db

import (
	"os"

	// This blank import is necessary to have the driver
	_ "github.com/mattn/go-sqlite3"
)

func newSqliteConfig() DBConfiger {
	return &SqLiteConfig{}
}

type SqLiteConfig struct {
}

func (c *SqLiteConfig) getConnectionString() string {
	return os.Getenv(envdbDatabase)
}

func (c *SqLiteConfig) getConnectionName() string {
	return DriverNameSqLite
}
