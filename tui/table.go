package tui

import (
	"keepass-tui/db"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var Table *tview.Table

func newTable() *tview.Table {
	Table = tview.NewTable().
		SetSelectable(true, true).
		SetFixed(1, 6)

	showAllRows()

	return Table
}

func newHeaderCell(text string) *tview.TableCell {
	return tview.NewTableCell(text).SetBackgroundColor(tcell.ColorGray).SetExpansion(1)
}

func newDataCell(text string, isPwd bool) *tview.TableCell {
	var cell *tview.TableCell

	if isPwd {
		cell = tview.NewTableCell("*****")
	} else {
		cell = tview.NewTableCell(text)
	}

	return cell.SetClickedFunc(func() bool {
		clipboard.WriteAll(text)
		return false
	}).SetMaxWidth(25)
}

func showHeaderRow() {
	Table.SetCell(0, 0, newHeaderCell("#").SetExpansion(0))
	Table.SetCell(0, 1, newHeaderCell("title"))
	Table.SetCell(0, 2, newHeaderCell("username"))
	Table.SetCell(0, 3, newHeaderCell("password"))
	Table.SetCell(0, 4, newHeaderCell("url"))
	Table.SetCell(0, 5, newHeaderCell("notes"))
}

func showAllRows() {
	showHeaderRow()
	filterRows("")
}

func filterRows(filter string) {
	Table.Clear()
	showHeaderRow()

	count := 0
	for _, entry := range db.Data {

		if !strings.Contains(entry.Title, filter) &&
			!strings.Contains(entry.Username, filter) &&
			!strings.Contains(entry.Url, filter) &&
			!strings.Contains(entry.Notes, filter) {
			continue
		}

		showRow(entry, count)
		count++
	}
}

func showRow(entry db.Entry, i int) {
	Table.SetCell(i+1, 0, newDataCell(strconv.Itoa(i), false))
	Table.SetCell(i+1, 1, newDataCell(entry.Title, false))
	Table.SetCell(i+1, 2, newDataCell(entry.Username, false))
	Table.SetCell(i+1, 3, newDataCell(entry.Password, true))
	Table.SetCell(i+1, 4, newDataCell(entry.Url, false))
	Table.SetCell(i+1, 5, newDataCell(entry.Notes, false))
}
