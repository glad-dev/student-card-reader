package main

import (
	"embed"
	"fmt"
	"os"

	"rucksack/dir"
	"rucksack/log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	// Init dirs
	err := dir.Create()
	if err != nil {
		fmt.Printf("Failed to create dir: %s\n", err)

		os.Exit(1)
	}

	// Init Logs
	log.Init()
}

func main() {
	// Create an instance of the app structure
	app, err := NewApp()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)

		os.Exit(1)
	}

	log.Error("This is a message", "key1", "value1", "key2", "value2", "key3", "value3")

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "wails-test",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
