package tui

import (
	"dboost/postgresql"
	"dboost/ui"

	"github.com/rivo/tview"
)

type TUI struct {
	App *tview.Application
	// sql      *mysql.Mysql
	sql *postgresql.Postgresql
	ui  *ui.UI
}

func New() *TUI {
	// ds := "root:pass@(localhost:3306)/"
	ds := "postgres://postgres:pass@localhost:5432/dvdrental?sslmode=disable"

	t := &TUI{
		App: tview.NewApplication(),
		// sql:      mysql.New(ds),
		sql: postgresql.New(ds),
		ui:  ui.New(),
	}

	return t
}

func (t *TUI) Run() error {
	layout := t.ui.Draw()
	t.setData()
	t.setKeyEvent()
	t.setEvent()

	return t.App.SetRoot(layout, true).SetFocus(layout).EnableMouse(true).Run()
}

func (t *TUI) setData() error {
	dbs, err := t.sql.ListDBs()
	if err != nil {
		return err
	}

	t.ui.SetDBs(dbs)

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

func (t *TUI) query(query string) error {
	records, err := t.sql.CustomQuery(query)
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
