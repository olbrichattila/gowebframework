package db

type DBConfiger interface {
	getConnectionName() string
	getConnectionString() string
}
