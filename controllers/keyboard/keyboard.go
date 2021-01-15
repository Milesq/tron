package keyboard

import (
	"fmt"

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
		_, key, err := keyboard.GetKey()

		if err != nil {
			fmt.Println("Keyboard listener disabled")
			break
		}

		switch key {
		case keyboard.KeyEsc:
			game.Emit(event.Exit, -1)
			break

			// ██╗   ██╗██████╗
			// ██║   ██║██╔══██╗
			// ██║   ██║██████╔╝
			// ██║   ██║██╔═══╝
			// ╚██████╔╝██║
			//  ╚═════╝ ╚═╝
		case keyboard.KeyArrowUp:
			game.Emit(event.Up, 0)
		case keyboard.KeyCtrlW:
			game.Emit(event.Up, 1)

			// ██████╗ ██╗ ██████╗ ██╗  ██╗████████╗
			// ██╔══██╗██║██╔════╝ ██║  ██║╚══██╔══╝
			// ██████╔╝██║██║  ███╗███████║   ██║
			// ██╔══██╗██║██║   ██║██╔══██║   ██║
			// ██║  ██║██║╚██████╔╝██║  ██║   ██║
			// ╚═╝  ╚═╝╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝
		case keyboard.KeyArrowRight:
			game.Emit(event.Right, 0)
		case keyboard.KeyCtrlD:
			game.Emit(event.Right, 1)

			// ██████╗  ██████╗ ██╗    ██╗███╗   ██╗
			// ██╔══██╗██╔═══██╗██║    ██║████╗  ██║
			// ██║  ██║██║   ██║██║ █╗ ██║██╔██╗ ██║
			// ██║  ██║██║   ██║██║███╗██║██║╚██╗██║
			// ██████╔╝╚██████╔╝╚███╔███╔╝██║ ╚████║
			// ╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝
		case keyboard.KeyArrowDown:
			game.Emit(event.Down, 0)
		case keyboard.KeyCtrlS:
			game.Emit(event.Down, 1)

			// ██╗     ███████╗███████╗████████╗
			// ██║     ██╔════╝██╔════╝╚══██╔══╝
			// ██║     █████╗  █████╗     ██║
			// ██║     ██╔══╝  ██╔══╝     ██║
			// ███████╗███████╗██║        ██║
			// ╚══════╝╚══════╝╚═╝        ╚═╝
		case keyboard.KeyArrowLeft:
			game.Emit(event.Left, 0)
		case keyboard.KeyCtrlA:
			game.Emit(event.Left, 1)
		}
	}
}
