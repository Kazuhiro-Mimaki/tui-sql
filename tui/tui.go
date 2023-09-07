package tui

import (
	"github.com/rivo/tview"
)

type TUI struct {
	App *tview.Application

	Flex *tview.Flex

	tables *tview.List
}

func New() *TUI {
	t := &TUI{
		App:    tview.NewApplication(),
		Flex:   tview.NewFlex(),
		tables: tview.NewList(),
	}

	t.tables.
		ShowSecondaryText(false).
		SetTitle("Tables").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)
	t.setTables()

	t.Flex.
		AddItem(t.tables, 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Main"), 0, 4, false)

	return t
}

func (tui *TUI) Run() error {
	tui.setEventKey()
	return tui.App.SetRoot(tui.Flex, true).SetFocus(tui.Flex).Run()
}

func (tui *TUI) queueUpdateDraw(f func()) {
	go func() {
		tui.App.QueueUpdateDraw(f)
	}()
}

func (tui *TUI) setFocus(p tview.Primitive) {
	tui.queueUpdateDraw(func() {
		tui.App.SetFocus(p)
	})
}

func (tui *TUI) setTables() {
	tui.queueUpdateDraw(func() {
		tui.tables.Clear()
		for _, table := range []string{"table1", "table2", "table3"} {
			tui.tables.AddItem(table, "", 0, nil)
		}
	})
}
