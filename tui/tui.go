package tui

import (
	"dboost/mysql"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	App       *tview.Application
	sql       *mysql.Mysql
	flex      *tview.Flex
	dbDD      *tview.DropDown
	tableList *tview.List
	records   *tview.Table
}

func New() *TUI {
	ds := "root:pass@(localhost:3306)/"

	t := &TUI{
		App:       tview.NewApplication(),
		sql:       mysql.New(ds),
		flex:      tview.NewFlex(),
		dbDD:      tview.NewDropDown(),
		tableList: tview.NewList(),
		records:   tview.NewTable(),
	}

	return t
}

func (tui *TUI) Run() error {
	tui.drawLayout()
	tui.setData()
	tui.setKeyEvent()
	tui.setEvent()
	return tui.App.SetRoot(tui.flex, true).SetFocus(tui.flex).Run()
}

func (tui *TUI) drawLayout() {
	tui.tableList.
		ShowSecondaryText(false).
		SetTitle("Tables").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	tui.dbDD.
		SetTitle("Database").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	side := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tui.dbDD, 0, 1, true).
		AddItem(tui.tableList, 0, 5, false)

	tui.records.
		Select(0, 0).
		SetFixed(1, 0).
		SetSelectable(true, true).
		SetTitle("Records").
		SetBorder(true)

	tui.flex.
		AddItem(side, 0, 1, false).
		AddItem(tui.records, 0, 4, false)
}

func (tui *TUI) setData() error {
	dbs, err := tui.sql.ListDBs()
	if err != nil {
		return err
	}
	tui.setDBs(dbs)

	tables, err := tui.sql.ListTables(dbs[0])
	if err != nil {
		return err
	}
	tui.setTables(tables)

	records, err := tui.sql.ListRecords(tables[0])
	if err != nil {
		return err
	}
	tui.setRecords(records)

	return nil
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

func (tui *TUI) setTables(tables []string) {
	tui.queueUpdateDraw(func() {
		tui.tableList.Clear()
		for _, table := range tables {
			tui.tableList.AddItem(table, "", 0, nil)
		}
	})
}

func (tui *TUI) setDBs(dbs []string) {
	tui.queueUpdateDraw(func() {
		tui.dbDD.SetOptions(dbs, nil)
		tui.dbDD.SetCurrentOption(0)
	})
}

func (tui *TUI) setRecords(records [][]*string) {
	tui.queueUpdateDraw(func() {
		tui.records.Clear().ScrollToBeginning()

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

				tui.records.SetCell(
					i, j,
					&tview.TableCell{
						Text:          cellValue,
						Color:         cellColor,
						NotSelectable: notSelectable,
					},
				)
			}
		}
	})
}

func (tui *TUI) selectDB(db string) error {
	tables, err := tui.sql.ListTables(db)
	if err != nil {
		return err
	}
	tui.setTables(tables)
	return nil
}

func (tui *TUI) selectTable(table string) error {
	records, err := tui.sql.ListRecords(table)
	if err != nil {
		return err
	}
	tui.setRecords(records)
	return nil
}
