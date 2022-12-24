package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var Conn *sql.DB

func Setup() {
	conn, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	Conn = conn
}

func Close() {
	Conn.Close()
	fmt.Println("database connection is closed.")
}
