package tui

import (
	"dboost/mysql"
	"dboost/ui"

	"github.com/rivo/tview"
)

type TUI struct {
	App *tview.Application
	sql *mysql.Mysql
	ui  *ui.UI
}

func New() *TUI {
	ds := "root:pass@(localhost:3306)/"

	t := &TUI{
		App: tview.NewApplication(),
		sql: mysql.New(ds),
		ui:  ui.New(),
	}

	return t
}

func (t *TUI) Run() error {
	layout := t.ui.Draw()
	t.setData()
	t.setKeyEvent()
	t.setEvent()
	return t.App.SetRoot(layout, true).SetFocus(layout).Run()
}

func (t *TUI) setData() error {
	dbs, err := t.sql.ListDBs()
	if err != nil {
		return err
	}

	tables, err := t.sql.ListTables(dbs[0])
	if err != nil {
		return err
	}

	records, err := t.sql.ListRecords(tables[0])
	if err != nil {
		return err
	}

	t.ui.SetDBs(dbs)
	t.ui.SetTables(tables)
	t.ui.SetRecords(records)

	return nil
}

func (t *TUI) selectDB(db string) error {
	tables, err := t.sql.ListTables(db)
	if err != nil {
		return err
	}
	t.ui.SetTables(tables)
	return nil
}

func (t *TUI) selectTable(table string) error {
	records, err := t.sql.ListRecords(table)
	if err != nil {
		return err
	}
	t.ui.SetRecords(records)
	return nil
}

func (t *TUI) queueUpdateDraw(f func()) {
	go func() {
		t.App.QueueUpdateDraw(f)
	}()
}
