package database

import (
	"database/sql"
	"fmt"
	"time"

	"rucksack/log"

	_ "github.com/mattn/go-sqlite3"
)

func allowedToGiveOutBags(db *sql.DB, maxBags int) (bool, error) {
	count, err := getDistributedBagCount(db)
	if err != nil {
		return false, err
	}

	return count < maxBags, nil
}

func getDistributedBagCount(db *sql.DB) (int, error) {
	var count int
	t := time.Now()
	err := db.QueryRow(fmt.Sprintf(`
		SELECT COUNT(ts) FROM rucksack
		WHERE
    	-- Check if it currently is after 12:15 (end of first shift)
    	CASE
        	WHEN '%s' < '12:15' THEN
				-- It is before 12:15
				ts BETWEEN
					datetime(date('now'), '00:00:00') AND
					datetime(date('now'), '12:15:00')
        	ELSE
				-- It is after 12:15
				ts BETWEEN
					datetime(date('now'), '12:15:01') AND
					datetime(date('now'), '23:59:59')
    	END;
	`, t.Format("15:04"))).Scan(&count)
	if err != nil {
		log.Error("Failed to query database for rucksack count", "error", err)

		return -1, fmt.Errorf("could not count bags: %w", err)
	}

	return count, nil
}
