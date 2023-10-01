package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Postgresql struct {
	db *sql.DB
}

func New(dataSource string) *Postgresql {
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatalf("failed to sql.Open: %v\n", err)
	}

	return &Postgresql{db}
}

func (m *Postgresql) Ping() error {
	err := m.db.Ping()
	if err != nil {
		return fmt.Errorf("failed to db.Ping: %v\n", err)
	}
	return nil
}

func (m *Postgresql) Close() error {
	err := m.db.Close()
	if err != nil {
		return fmt.Errorf("failed to db.Close: %v\n", err)
	}
	return nil
}

func (m *Postgresql) ListDBs() ([]string, error) {
	rows, err := m.db.Query("SELECT datname FROM pg_database")
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

func (m *Postgresql) ListTables(db string) ([]string, error) {
	rows, err := m.db.Query(fmt.Sprintf("SELECT table_name FROM information_schema.tables WHERE table_catalog = '%s';", db))
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

func (m *Postgresql) ListRecords(table string) (data [][]*string, err error) {
	rows, err := m.db.Query(fmt.Sprintf("SELECT * FROM %s", table))
	if err != nil {
		return nil, err
	}

	data, err = scanRows(rows)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func scanRows(rows *sql.Rows) (data [][]*string, err error) {
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

func (m *Postgresql) CustomQuery(query string) (data [][]*string, err error) {
	rows, err := m.db.Query(fmt.Sprintf("%s", query))
	if err != nil {
		return nil, err
	}

	data, err = scanRows(rows)
	if err != nil {
		return nil, err
	}

	return data, nil
}
