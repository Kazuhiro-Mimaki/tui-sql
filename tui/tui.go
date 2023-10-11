package tui

import (
	"errors"

	"dboost/config"
	"dboost/ds"
	"dboost/ds/db/mysql"
	"dboost/ds/db/postgresql"

	"dboost/ui"

	"github.com/rivo/tview"
)

type TUI struct {
	App    *tview.Application
	ds     ds.DataSource
	ui     *ui.UI
	config *config.Config
}

func New() *TUI {
	t := &TUI{
		App:    tview.NewApplication(),
		ui:     ui.New(),
		config: config.New(),
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

func (t *TUI) setData() {
	dbs := t.config.ListDBs()
	t.ui.SetDBs(dbs)
}

func (t *TUI) selectDB(db string) error {
	ds, err := t.getDataSource(db)
	if err != nil {
		return err
	}
	t.ds = ds

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

func (t *TUI) getDataSource(db string) (ds.DataSource, error) {
	conn, err := t.config.GetConn(db)
	if err != nil {
		return nil, err
	}

	switch conn.Type {
	case "mysql":
		return mysql.New(conn.Dsn), nil
	case "postgres":
		return postgresql.New(conn.Dsn), nil
	default:
		return nil, errors.New("invalid connection type")
	}
}
