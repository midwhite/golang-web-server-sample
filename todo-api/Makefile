start:
	air -c .air.toml
build:
	go build -o cmd/api/run cmd/api/main.go
init-db:
	psql -U postgres -h todo-api-db -c "CREATE DATABASE todo_api_development;"
migrate:
	migrate -source file://migration -database ${DATABASE_URL} up
psql:
	psql -U postgres -h todo-api-db todo_api_development
