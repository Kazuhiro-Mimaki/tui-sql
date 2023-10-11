package ui

import (
	"github.com/rivo/tview"
)

type UI struct {
	DBDD      *tview.DropDown
	TableList *tview.List
	Query     *tview.InputField
	Preview   *Preview
}

func New() *UI {
	t := &UI{
		DBDD:      tview.NewDropDown(),
		TableList: tview.NewList(),
		Query:     tview.NewInputField(),
		Preview:   NewPreview(),
	}

	return t
}

func (ui *UI) Draw() *tview.Flex {
	ui.DBDD.
		SetTitle("Database (Ctrl-A)").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	ui.TableList.
		ShowSecondaryText(false).
		SetTitle("Tables (Ctrl-F)").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	side := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(ui.DBDD, 0, 1, true).
		AddItem(ui.TableList, 0, 13, false)

	ui.Query.
		SetTitle("Query (Ctrl-E)").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	main := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(ui.Query, 0, 1, false).
		AddItem(ui.Preview.View, 0, 13, false)

	return tview.NewFlex().
		AddItem(side, 0, 1, false).
		AddItem(main, 0, 4, false)
}

func (ui *UI) SetDBs(dbs []string) {
	ui.DBDD.SetOptions(dbs, nil)
	ui.DBDD.SetCurrentOption(0)
}

func (ui *UI) SetTables(tables []string) {
	ui.TableList.Clear()
	for _, table := range tables {
		ui.TableList.AddItem(table, "", 0, nil)
	}
}
