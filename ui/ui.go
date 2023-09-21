package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UI struct {
	DBDD      *tview.DropDown
	TableList *tview.List
	Records   *tview.Table
}

func New() *UI {
	t := &UI{
		DBDD:      tview.NewDropDown(),
		TableList: tview.NewList(),
		Records:   tview.NewTable(),
	}

	return t
}

func (ui *UI) Draw() *tview.Flex {
	ui.TableList.
		ShowSecondaryText(false).
		SetTitle("Tables").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	ui.DBDD.
		SetTitle("Database").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	side := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(ui.DBDD, 0, 1, true).
		AddItem(ui.TableList, 0, 5, false)

	ui.Records.
		Select(0, 0).
		SetFixed(1, 0).
		SetSelectable(true, true).
		SetTitle("Records").
		SetBorder(true)

	return tview.NewFlex().
		AddItem(side, 0, 1, false).
		AddItem(ui.Records, 0, 4, false)
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

func (ui *UI) SetRecords(records [][]*string) {
	ui.Records.Clear().ScrollToBeginning()

	for i, row := range records {
		for j, col := range row {
			var cellValue string
			cellColor := tcell.ColorWhite
			notSelectable := false

			if col != nil {
				cellValue = *col
			}

			// カラム名の色はレコードと異なるものを指定
			if i == 0 {
				cellColor = tcell.ColorNavy
			}

			ui.Records.SetCell(
				i, j,
				&tview.TableCell{
					Text:          cellValue,
					Color:         cellColor,
					NotSelectable: notSelectable,
				},
			)
		}
	}
}
