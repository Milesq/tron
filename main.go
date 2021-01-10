package main

import (
	"time"

	"github.com/milesq/tron/tron"
)

func main() {
	game := tron.NewGame(tron.Config{
		Players:     []string{"red", "blue"},
		PlayerSpeed: 20,
		Size:        [2]int{100, 100},
	})

	for {
		game.Next()
		time.Sleep(time.Second / 60)
	}
}
