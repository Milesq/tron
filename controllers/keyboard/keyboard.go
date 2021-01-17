package keyboard

import (
	"fmt"

	tm "github.com/buger/goterm"
	"github.com/eiannone/keyboard"

	"github.com/milesq/tron/tron"
	"github.com/milesq/tron/tron/event"
)

// Controller .
func Controller(game *tron.Game) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()

		if err != nil {
			fmt.Println("Keyboard listener disabled")
			break
		}

		switch key {
		case keyboard.KeyEsc:
			game.Emit(event.Exit, -1)
			break

		case keyboard.KeyArrowUp:
			game.Emit(event.Up, tm.RED)

		case keyboard.KeyArrowRight:
			game.Emit(event.Right, tm.RED)

		case keyboard.KeyArrowDown:
			game.Emit(event.Down, tm.RED)

		case keyboard.KeyArrowLeft:
			game.Emit(event.Left, tm.RED)
		}

		switch char {
		case 'W', 'w':
			game.Emit(event.Up, tm.BLUE)
		case 'S', 's':
			game.Emit(event.Down, tm.BLUE)
		case 'A', 'a':
			game.Emit(event.Left, tm.BLUE)
		case 'D', 'd':
			game.Emit(event.Right, tm.BLUE)

		case 'I', 'i':
			game.Emit(event.Up, tm.YELLOW)
		case 'K', 'k':
			game.Emit(event.Down, tm.YELLOW)
		case 'J', 'j':
			game.Emit(event.Left, tm.YELLOW)
		case 'L', 'l':
			game.Emit(event.Right, tm.YELLOW)
		}
	}
}
