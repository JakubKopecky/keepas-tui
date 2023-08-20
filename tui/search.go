package tui

import (
	"crypto/sha256"
	"encoding/hex"
	"keepass-tui/db"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func newSearchBar() *tview.InputField {
	var searchbar *tview.InputField

	searchbar = tview.NewInputField().
		SetChangedFunc(func(text string) {
			filterRows(text)
		}).SetDoneFunc(func(key tcell.Key) {
		if searchbar.GetText()[0] == ':' {
			unlockDb(searchbar.GetText()[1:])
			showAllRows()
			searchbar.SetText("")
		}
	})
	return searchbar
}

func unlockDb(pwd string) {
	hash := sha256.New()
	hash.Write([]byte(pwd))
	niceHash := hex.EncodeToString(hash.Sum(nil))

	db.OpenAndLoadDB(dbPath, niceHash)
}
