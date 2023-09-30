package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (t *TUI) setFocus(p tview.Primitive) {
	t.App.SetFocus(p)
}

func (t *TUI) setKeyEvent() {
	t.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			t.setFocus(t.ui.TableList)
		case tcell.KeyCtrlS:
			t.setFocus(t.ui.DBDD)
		case tcell.KeyCtrlR:
			t.setFocus(t.ui.Records)
		}
		return event
	})
}

func (t *TUI) setEvent() {
	t.ui.DBDD.SetSelectedFunc(func(db string, _ int) {
		t.selectDB(db)
	})

	t.ui.TableList.SetSelectedFunc(func(_ int, table, _ string, _ rune) {
		t.selectTable(table)
	})

	t.ui.Query.SetDoneFunc(func(key tcell.Key) {
		t.query(t.ui.Query.GetText())
	})
}
