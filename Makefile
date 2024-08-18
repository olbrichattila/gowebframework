db-recreate:
	@if [ -f ./database/database.sqlite ]; then rm ./database/database.sqlite; fi
	csvimporter data.csv vehicles ";"
	migrator migrate