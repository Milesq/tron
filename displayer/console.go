package displayer

import (
	"fmt"

	"github.com/milesq/tron/tron"
)

// ConsoleDisplayer .
func ConsoleDisplayer(state tron.GameState) {
	fmt.Println(state)
}
