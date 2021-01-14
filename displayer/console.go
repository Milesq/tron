package displayer

import (
	tm "github.com/buger/goterm"

	"github.com/milesq/tron/tron"
)

// ConsoleDisplayer .
func ConsoleDisplayer(state tron.GameState) {
	box := tm.NewBox(state.MapSize[0]|tm.PCT, state.MapSize[1], 0)

	for _, trace := range state.Players {
		currentPoint := trace[len(trace)-1]

		tm.MoveCursor(currentPoint.X, currentPoint.Y)
		tm.Print(tm.Color("@", tm.RED))
	}

	tm.Print(tm.MoveTo(box.String(), 1|tm.PCT, 10|tm.PCT))
}
