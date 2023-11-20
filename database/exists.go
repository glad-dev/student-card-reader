package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"rucksack/log"

	_ "github.com/mattn/go-sqlite3"
)

func existingEntry(db *sql.DB, id string, debug bool) (string, error) {
	if debug && id == "daca4ceca913a9e07888469dc47c5bee5caacebb32000914230b39787549ac19" { // The hash of my card, added for debug purposes
		return "", nil
	}

	var timestamp string
	err := db.QueryRow("SELECT ts FROM rucksack WHERE id_hash = ? LIMIT 1", id).Scan(&timestamp)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Error("Failed to query database for existing record: %s", err)

		return "", fmt.Errorf("could not query db: %w", err)
	}

	if len(timestamp) > 0 {
		ts, err := time.Parse("2006-01-02T15:04:05Z", timestamp)
		if err != nil {
			log.Error("Failed to parse timestamp", "timestamp", timestamp, "error", err)
		} else {
			timestamp = ts.Format("02.01.2006 15:04")
		}
	}

	return timestamp, nil
}
