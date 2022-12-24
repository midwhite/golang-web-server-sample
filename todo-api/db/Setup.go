package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

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

func readMigrationFiles(pattern string) []string {
	filepaths, _ := filepath.Glob("/app/db/migrations/" + pattern)
	result := make([]string, len(filepaths))

	for i, path := range filepaths {
		file, _ := os.Open(path)
		stat, _ := file.Stat()
		sql := make([]byte, stat.Size())
		file.Read(sql)
		result[i] = string(sql)
	}

	return result
}

func Migrate() {
	sqls := readMigrationFiles("*.up.sql")

	for _, sql := range sqls {
		_, err := Conn.Exec(sql)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func Rollback() {
	sqls := readMigrationFiles("*.down.sql")

	for _, sql := range sqls {
		_, err := Conn.Exec(sql)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}

func Close() {
	Conn.Close()
	fmt.Println("database connection is closed.")
}
