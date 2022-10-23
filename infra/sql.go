package infra

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("sqlite3", "db/diaries.db")
	if err != nil {
		panic(err)
	}
	tableName := "diaries"
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title STRING NOT NULL,
			content STRING NOT NULL,
			tag INT NOT NULL,
			created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
			updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
		)`, tableName)
	_, err = conn.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
	sqlHandler := &SqlHandler{conn}
	return sqlHandler
}
