package main

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
	"github.com/milesq/tron/displayer"
	"github.com/milesq/tron/tron"
)

func main() {
	tm.Clear()

	game := tron.NewGame(tron.Config{
		Players:     []string{"red", "blue"},
		PlayerSpeed: 20,
		Size:        [2]int{100, 100},
	})

	quit := make(chan int)

	go game.UpdateWithInterval(time.Second/30, quit)

	go func() {
		displayInterval := time.NewTicker(time.Second / 30)

		for range displayInterval.C {
			displayer.ConsoleDisplayer(game.State)
		}
	}()

	score := <-quit

	fmt.Println("score: ", score)
}
