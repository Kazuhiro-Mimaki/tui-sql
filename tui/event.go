package tui

import "github.com/gdamore/tcell/v2"

func (tui *TUI) setKeyEvent() {
	tui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			tui.setFocus(tui.tableList)
		case tcell.KeyCtrlS:
			tui.setFocus(tui.dbDD)
		case tcell.KeyCtrlR:
			tui.setFocus(tui.records)
		}
		return event
	})
}

func (tui *TUI) setEvent() {
	tui.queueUpdateDraw(func() {
		tui.dbDD.SetSelectedFunc(func(db string, _ int) {
			tui.selectDB(db)
		})

		tui.tableList.SetSelectedFunc(func(_ int, table, _ string, _ rune) {
			tui.selectTable(table)
		})
	})
}
