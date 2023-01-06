start:
	air -c .air.toml
test:
	go test ./...
build:
	go build -o cmd/api/run cmd/api/main.go
proto:
	protoc -I . ./user-service.proto --go_out=./proto/
generate-models:
	sqlboiler psql --wipe
init-db:
	psql -U postgres -h todo-api-db -c "CREATE DATABASE todo_api_development;"
	psql -U postgres -h todo-api-db -c "CREATE DATABASE todo_api_test;"
migrate:
	migrate -source file://db/migrations -database ${DATABASE_URL} up
rollback:
	migrate -source file://db/migrations -database ${DATABASE_URL} down
psql:
	psql -U postgres -h todo-api-db todo_api_development
