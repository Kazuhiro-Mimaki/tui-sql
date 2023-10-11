package main

import "tui-sql/tui"

func main() {
	t := tui.New()

	if err := t.Run(); err != nil {
		panic(err)
	}
}
