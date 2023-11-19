package database

import (
	"database/sql"
	"fmt"
	"path"
	"rucksack/log"

	"rucksack/dir"

	_ "github.com/mattn/go-sqlite3"
)

func OpenOrInitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path.Join(dir.Get(), "database.db"))
	if err != nil {
		log.Error("Failed to open database", "error", err)

		return nil, fmt.Errorf("could not open db: %w", err)
	}

	// Create table if it does not exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS rucksack (id_hash text PRIMARY KEY, ts timestamp NOT NULL)")
	if err != nil {
		log.Error("Failed to create table", "error", err)

		return nil, fmt.Errorf("could not create table: %w", err)
	}

	return db, nil
}
