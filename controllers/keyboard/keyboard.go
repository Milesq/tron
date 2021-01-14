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
		char, key, err := keyboard.GetKey()

		if err != nil {
			fmt.Println("Keyboard listener disabled")
			break
		}

		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)

		if key == keyboard.KeyEsc {
			game.Emit(event.Exit)
			break
		}
	}
}
