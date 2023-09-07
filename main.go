package main

import (
	"dboost/tui"
)

func main() {
	t := tui.New()

	if err := t.Run(); err != nil {
		panic(err)
	}
}
