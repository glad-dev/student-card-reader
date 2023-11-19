package database

import (
	"database/sql"
	"fmt"
	"time"

	"rucksack/log"

	_ "github.com/mattn/go-sqlite3"
)

func InsertStudentID(db *sql.DB, id string, debug bool) error {
	b, err := allowedToGiveOutBags(db, 300)
	if b {
		log.Error("Already distributed allowed bag count")

		return fmt.Errorf("not allowed to give out bags")
	} else if err != nil {
		log.Error("Failed to check if allowed to give out bags", "error", err)

		return fmt.Errorf("could not check count: %w", err)
	}

	ts, err := existingEntry(db, id, debug)
	if err != nil {
		log.Error("Failed to check for existing entry", "error", err, "id", id)

		return fmt.Errorf("could not check for existing entries: %w", err)
	} else if len(ts) > 0 {
		log.Error("Existing entry", "id", id, "time-stamp", ts)

		return fmt.Errorf("student already got a bag on %s", ts)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Error("Failed to begin transaction", "error", err)

		return fmt.Errorf("could not begin transaction: %w", err)
	}

	stmt, err := tx.Prepare("INSERT INTO rucksack (id_hash, ts) VALUES (?, ?)")
	if err != nil {
		log.Error("Failed to prepare statement", "error", err)

		return fmt.Errorf("could not prepare statmement: %w", err)
	}

	_, err = stmt.Exec(id, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Error("Failed to execute statement", "error", err)

		return fmt.Errorf("could not execute statement: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Error("Failed to commit transaction", "error", err)

		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}