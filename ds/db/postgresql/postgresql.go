package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	"dboost/ds"

	_ "github.com/lib/pq"
)

type Postgresql struct {
	db *sql.DB
}

func New(dsn string) ds.DataSource {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to sql.Open: %v\n", err)
	}

	return &Postgresql{db}
}

func (p *Postgresql) Ping() error {
	err := p.db.Ping()
	if err != nil {
		return fmt.Errorf("failed to db.Ping: %v\n", err)
	}
	return nil
}

func (p *Postgresql) Close() error {
	err := p.db.Close()
	if err != nil {
		return fmt.Errorf("failed to db.Close: %v\n", err)
	}
	return nil
}

func (p *Postgresql) ListDBs() ([]string, error) {
	rows, err := p.db.Query("SELECT datname FROM pg_database")
	if err != nil {
		return nil, err
	}

	dbs := []string{}

	for rows.Next() {
		var dbName string
		err = rows.Scan(&dbName)
		if err == nil {
			dbs = append(dbs, dbName)
		}
	}

	return dbs, err
}

func (p *Postgresql) ListTables(db string) ([]string, error) {
	rows, err := p.db.Query(fmt.Sprintf("SELECT table_name FROM information_schema.tables WHERE table_catalog = '%s';", db))
	if err != nil {
		return nil, err
	}

	tables := []string{}

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err == nil {
			tables = append(tables, tableName)
		}
	}

	return tables, err
}

func (p *Postgresql) ListRecords(table string) (data [][]*string, err error) {
	rows, err := p.db.Query(fmt.Sprintf("SELECT * FROM %s", table))
	if err != nil {
		return nil, err
	}

	data, err = p.scanRows(rows)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *Postgresql) scanRows(rows *sql.Rows) (data [][]*string, err error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var colNames []*string
	for _, col := range cols {
		colName := col
		colNames = append(colNames, &colName)
	}

	data = [][]*string{}

	// カラム名を最初に設定
	data = append(data, colNames)

	for rows.Next() {
		row := make([]*string, len(cols))
		rowPointers := make([]interface{}, len(cols))
		for i := range row {
			rowPointers[i] = &row[i]
		}

		err = rows.Scan(rowPointers...)
		if err != nil {
			return nil, err
		}

		data = append(data, row)
	}

	return data, err
}

func (p *Postgresql) CustomQuery(query string) (data [][]*string, err error) {
	rows, err := p.db.Query(fmt.Sprintf("%s", query))
	if err != nil {
		return nil, err
	}

	data, err = p.scanRows(rows)
	if err != nil {
		return nil, err
	}

	return data, nil
}
