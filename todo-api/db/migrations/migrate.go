package migrations

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
)

func Migrate(conn *pgx.Conn) {
	filenames, _ := filepath.Glob("./db/migrations/sql/*.sql")

	for _, filename := range filenames {
		fmt.Println("migration file detected: " + filename)

		file, _ := os.Open(filename)
		stat, _ := file.Stat()
		sql := make([]byte, stat.Size())

		file.Read(sql)

		_, err := conn.Exec(context.Background(), string(sql))

		if err != nil {
			fmt.Fprintf(os.Stderr, "migration is failed: %v", err)
			os.Exit(1)
		}
	}

	fmt.Println("migration finished successfully.")
}
