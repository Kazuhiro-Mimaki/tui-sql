package tui

import (
	"dboost/mysql"

	"github.com/rivo/tview"
)

type TUI struct {
	App       *tview.Application
	sql       *mysql.Mysql
	flex      *tview.Flex
	dbDD      *tview.DropDown
	tableList *tview.List
}

func New() *TUI {
	ds := "root:pass@(localhost:3306)/"

	t := &TUI{
		App:       tview.NewApplication(),
		sql:       mysql.New(ds),
		flex:      tview.NewFlex(),
		dbDD:      tview.NewDropDown(),
		tableList: tview.NewList(),
	}

	return t
}

func (tui *TUI) Run() error {
	tui.drawLayout()
	tui.setData()
	tui.setEventKey()
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

	tui.flex.
		AddItem(side, 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Main"), 0, 4, false)
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
		tui.dbDD.SetOptions(dbs, func(db string, index int) {
			tui.selectDB(db)
		})
		tui.dbDD.SetCurrentOption(0)
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
