package mysql

import (
	"database/sql"
	"fmt"
	"log"
)

type Mysql struct {
	db *sql.DB
}

func New(dataSource string) *Mysql {
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("failed to sql.Open: %v\n", err)
	}

	return &Mysql{db}
}

func (m *Mysql) Ping() error {
	err := m.db.Ping()
	if err != nil {
		return fmt.Errorf("failed to db.Ping: %v\n", err)
	}
	return nil
}

func (m *Mysql) Close() error {
	err := m.db.Close()
	if err != nil {
		return fmt.Errorf("failed to db.Close: %v\n", err)
	}
	return nil
}

func (m *Mysql) ListDBs() ([]string, error) {
	rows, err := m.db.Query("SHOW DATABASES")
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

func (m *Mysql) ListTables(db string) ([]string, error) {
	_, err := m.db.Exec("USE " + db)
	if err != nil {
		return nil, err
	}

	rows, err := m.db.Query("SHOW TABLES")
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
