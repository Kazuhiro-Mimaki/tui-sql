package main

import (
	"log"

	"dboost/mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ds := "test_user:password@tcp(127.0.0.1:3306)/world"
	m := mysql.New(ds)
	err := m.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// t := tui.New()

	// if err := t.Run(); err != nil {
	// 	panic(err)
	// }
}
