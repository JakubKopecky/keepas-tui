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
			if len(text) != 0 && text[0] == ':' {
				searchbar.SetFieldBackgroundColor(tcell.ColorWhite)
			} else {
				searchbar.SetFieldBackgroundColor(tcell.ColorBlue)
			}
			filterRows(text)
		}).SetDoneFunc(func(key tcell.Key) {
		if searchbar.GetText()[0] == ':' {
			unlockDb(searchbar.GetText()[1:])
			searchbar.SetText("")
			showAllRows()
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
