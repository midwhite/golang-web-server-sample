package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/midwhite/golang-web-server-sample/todo-api/db/migrations"
)

var Conn *pgx.Conn

func Setup() {
	databaseURL := "postgres://postgres@todo-api-db:5432/todo_api_development"
	conn, err := pgx.Connect(context.Background(), databaseURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	migrations.Migrate(conn)

	Conn = conn
}

func Close() {
	Conn.Close(context.Background())
	fmt.Println("database connection is closed.")
}
