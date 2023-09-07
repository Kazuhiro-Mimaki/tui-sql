package tui

import "github.com/gdamore/tcell/v2"

func (tui *TUI) setEventKey() {
	tui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			tui.setFocus(tui.tables)
		}
		return event
	})
}
