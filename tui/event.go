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
			t.setFocus(t.ui.DBDD)
		case tcell.KeyCtrlF:
			t.setFocus(t.ui.TableList)
		case tcell.KeyCtrlE:
			t.setFocus(t.ui.Query)
		case tcell.KeyCtrlR:
			t.ui.Preview.SwitchPage("records")
			t.setFocus(t.ui.Preview.Records)
		case tcell.KeyCtrlS:
			t.ui.Preview.SwitchPage("schemas")
			t.setFocus(t.ui.Preview.Schemas)
		}
		return event
	})
}

func (t *TUI) setEvent() {
	t.ui.DBDD.SetSelectedFunc(func(db string, _ int) {
		t.selectDB(db)
		t.setFocus(t.ui.TableList)
	})

	t.ui.TableList.SetSelectedFunc(func(_ int, table, _ string, _ rune) {
		t.selectTable(table)
		t.setFocus(t.ui.Preview.Records)
	})

	t.ui.Query.SetDoneFunc(func(key tcell.Key) {
		t.query(t.ui.Query.GetText())
		t.setFocus(t.ui.Preview.Records)
	})
}
