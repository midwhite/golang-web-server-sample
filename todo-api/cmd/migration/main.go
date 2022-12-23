package main

import (
	"github.com/midwhite/golang-web-server-sample/todo-api/db"
	"github.com/midwhite/golang-web-server-sample/todo-api/db/migrations"
)

func main() {
	db.Setup()
	defer db.Close()

	migrations.Migrate(db.Conn)
}
