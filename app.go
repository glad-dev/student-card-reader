package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"

	"rucksack/database"
	"rucksack/log"
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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) AddStudent(id string) error {
	tmp := sha256.Sum256([]byte(id))
	id = hex.EncodeToString(tmp[:])

	return database.InsertStudentID(a.db, id, a.debug)
}
