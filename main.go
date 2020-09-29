package main

import (
	"chill/app"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "World!"
}

func main() {
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	a := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "chill",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	// work 30 minutes then chill for 15 secondss
	a.Bind(app.NewTimer(30*60, 15))
	a.Bind(app.BuildQuotes())
	a.Bind(app.NewTasks())
	a.Bind(basic)
	a.Run()
}
