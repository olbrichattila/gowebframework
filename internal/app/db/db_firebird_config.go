package db

import (
	"fmt"
	"os"

	// This blank import is necessary to have the driver
	_ "github.com/nakagami/firebirdsql"
)

type FirebirdConfig struct {
}

func newFirebirdConfig() DBConfiger {
	return &FirebirdConfig{}
}

func (c *FirebirdConfig) getConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@%s:%s%s",
		os.Getenv(envdbUserName),
		os.Getenv(envdbPassword),
		os.Getenv(envdbHost),
		os.Getenv(envdbPort),
		os.Getenv(envdbDatabase),
	)
}

func (c *FirebirdConfig) getConnectionName() string {
	return DriverNameFirebird
}
