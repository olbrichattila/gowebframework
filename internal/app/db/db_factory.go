package db

import (
	"fmt"
	"os"
)

func NewDBFactory() DBFactoryer {
	return &DBconf{}
}

type DBFactoryer interface {
	GetConnectionConfig() (DBConfiger, error)
}

const (
	driverNameFirebird = "firebirdsql"
	driverNameSqLite   = "sqlite3"
	driverNameMySQL    = "mysql"
	driverNamePostgres = "postgres"

	dbConnectionTypeSqLite   = "sqlite"
	dbConnectionTypeMySQL    = "mysql"
	dbConnectionTypePgSQL    = "pgsql"
	dbConnectionTypeFirebird = "firebird"
	dbConnectionTypeMemory   = "memory"

	envdbConnection = "DB_CONNECTION"
	envdbUserName   = "DB_USERNAME"
	envdbPassword   = "DB_PASSWORD"
	envdbHost       = "DB_HOST"
	envdbPort       = "DB_PORT"
	envdbDatabase   = "DB_DATABASE"
	envdbSSLMode    = "DB_SSLMODE"
)

type DBconf struct {
	config DBConfiger
}

func (c *DBconf) GetConnectionConfig() (DBConfiger, error) {
	dbConnection := os.Getenv(envdbConnection)

	switch dbConnection {
	case dbConnectionTypeSqLite:
		return newSqliteConfig(), nil
	case dbConnectionTypeMySQL:
		return newMySQLConfig(), nil
	case dbConnectionTypePgSQL:
		return newPgsqlConfig(), nil
	case dbConnectionTypeFirebird:
		return newFirebirdConfig(), nil
	case dbConnectionTypeMemory:
		return newMemoryDBConfig(), nil
	default:
		return nil, fmt.Errorf("invalid DB_CONNECTION %s", dbConnection)
	}
}
