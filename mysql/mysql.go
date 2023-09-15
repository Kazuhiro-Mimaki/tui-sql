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
