db-recreate:
	@if [ -f ./database/database.sqlite ]; then rm ./database/database.sqlite; fi
	csvimporter data.csv vehicles ";"
	migrator migrate
switch-sqlite:
	cp .env.sqlite.example .env.migrator
	cp .env.sqlite.example .env.csvimporter
	cp .env.sqlite.example .env.csvexporter
switch-mysql:
	cp .env.mysql.example .env.migrator
	cp .env.mysql.example .env.csvimporter
	cp .env.mysql.example .env.csvexporter
switch-pgsql:
	cp .env.pgsql.example .env.migrator
	cp .env.pgsql.example .env.csvimporter
	cp .env.pgsql.example .env.csvexporter
switch-firebird:
	cp .env.firebird.example .env.migrator
	cp .env.firebird.example .env.csvimporter
	cp .env.firebird.example .env.csvexporter
lint:
	gocritic check ./...
	revive ./...
	golint ./...
	go vet ./...
	staticcheck ./...
	golangci-lint run
	goconst ./...
test:
	go test ./...
