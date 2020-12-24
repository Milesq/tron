package main

import (
	"time"

	tm "github.com/buger/goterm"
)

func main() {
	tm.Clear()

	for {
		tm.MoveCursor(1, 1)

		tm.Println("Current Time:", time.Now().Format(time.RFC1123))

		tm.Flush()

		time.Sleep(time.Second)
	}
}
