package main

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"

	keyCtrl "github.com/milesq/tron/controllers/keyboard"
	"github.com/milesq/tron/displayer"
	"github.com/milesq/tron/tron"
)

func main() {
	tm.Clear()

	game := tron.NewGame(tron.Config{
		Players:     []string{"red", "blue"},
		PlayerSpeed: 20,
		Size:        [2]int{150, 30},
		PlayerChar:  '@',
		BorderChar:  '-',
	})

	quit := make(chan int)

	go game.UpdateWithInterval(time.Second/30, quit)

	go func() {
		displayInterval := time.NewTicker(time.Second / 40)

		for range displayInterval.C {
			displayer.ConsoleDisplayer(game.State)
			tm.Flush()
		}
	}()

	keyCtrl.Controller(&game)

	score := <-quit

	fmt.Println("score: ", score)
}
