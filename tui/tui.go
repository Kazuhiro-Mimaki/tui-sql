package tui

import (
	"errors"
	"log"

	"dboost/ds"
	"dboost/mysql"
	"dboost/postgresql"
	"dboost/ui"

	"github.com/rivo/tview"
)

type TUI struct {
	App *tview.Application
	ds  ds.DataSource
	ui  *ui.UI
}

func New() *TUI {
	ds, err := getDataSource("mysql")
	if err != nil {
		log.Fatal(err)
	}

	t := &TUI{
		App: tview.NewApplication(),
		ds:  ds,
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
	dbs, err := t.ds.ListDBs()
	if err != nil {
		return err
	}

	t.ui.SetDBs(dbs)

	return nil
}

func (t *TUI) selectDB(db string) error {
	tables, err := t.ds.ListTables(db)
	if err != nil {
		return err
	}
	t.ui.SetTables(tables)
	return nil
}

func (t *TUI) selectTable(table string) error {
	records, err := t.ds.ListRecords(table)
	if err != nil {
		return err
	}
	t.ui.SetRecords(records)
	return nil
}

func (t *TUI) query(query string) error {
	records, err := t.ds.CustomQuery(query)
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

func getDataSource(connType string) (ds.DataSource, error) {
	switch connType {
	case "mysql":
		ds := "root:pass@(localhost:3306)/"
		return mysql.New(ds), nil
	case "postgres":
		ds := "postgres://postgres:pass@localhost:5432/dvdrental?sslmode=disable"
		return postgresql.New(ds), nil
	default:
		return nil, errors.New("invalid connection type")
	}
}
