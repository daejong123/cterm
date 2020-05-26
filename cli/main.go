package main

import (
	"github.com/daejong123/cterm"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	cterm.FetchData()
	cterm.Draw(app)
	cterm.Keybinds(app)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
