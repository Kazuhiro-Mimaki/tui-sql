package main

import (
	"dboost/tui"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	t := tui.New()

	if err := t.Run(); err != nil {
		panic(err)
	}
}
