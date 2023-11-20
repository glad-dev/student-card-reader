package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"rucksack/database"
	"rucksack/log"
	"strings"
)

// App struct
type App struct {
	ctx   context.Context
	db    *sql.DB
	debug bool
}

// NewApp creates a new App application struct
func NewApp() (*App, error) {
	db, err := database.OpenOrInitDB()
	if err != nil {
		return nil, err
	}

	return &App{
		db: db,
	}, nil
}

// startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	err := a.db.Close()
	log.Error("Failed to close db", "error", err)
}

func (a *App) AddStudent(id string, uni string) error {
	id = strings.TrimSuffix(id, "$")

	if !a.debug && uni == "TUM" && len(id) != 16 {
		return errors.New("invalid TUM/LMU ID. Try again")
	}

	tmp := sha256.Sum256([]byte(id))
	id = hex.EncodeToString(tmp[:])

	err := database.InsertStudentID(a.db, id, a.debug)
	if err != nil {
		return err
	}

	return nil
}
