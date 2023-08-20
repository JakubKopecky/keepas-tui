package tui

import (
	"github.com/rivo/tview"
)

var dbPath string

func CreateTui(path string) *tview.Application {
	dbPath = path

	app := tview.NewApplication()
	table := newTable()
	searchBar := newSearchBar()

	grid := tview.NewGrid().
		SetRows(0, 1).
		SetColumns(0).
		SetBorders(true).
		AddItem(table, 0, 0, 1, 1, 0, 0, false).
		AddItem(searchBar, 1, 0, 1, 1, 0, 0, true)

	app.SetRoot(grid, true).EnableMouse(true)

	return app
}
