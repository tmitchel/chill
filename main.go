package main

import (
	"chill/app"
	"os"
	"path/filepath"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	user, _ := os.UserHomeDir()
	db, err := app.Open(filepath.Join(user, "Documents", ".chill_storage.json"))
	if err != nil {
		return
	}

	a := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Title:     "chill",
		JS:        js,
		CSS:       css,
		Colour:    "#FFF",
		Resizable: true,
	})

	// work 30 minutes then chill for 15 secondss
	a.Bind(app.NewTimer(30*60, 15))
	a.Bind(app.BuildQuotes())
	a.Bind(app.NewTasks(db))
	a.Bind(app.NewStats(db))
	a.Run()
}
