package displayer

import (
	tm "github.com/buger/goterm"

	"github.com/milesq/tron/tron"
)

// ConsoleDisplayer .
func ConsoleDisplayer(state tron.GameState) {
	for player, trace := range state.Players {
		for _, pt := range trace {
			tm.MoveCursor(pt.X, pt.Y)
			tm.Print(tm.Color(string(state.BorderChar), player))
		}

		currentPoint := trace[len(trace)-1]
		tm.MoveCursor(currentPoint.X, currentPoint.Y)
		tm.Print(tm.Color(string(state.PlayerChar), player))
	}
}
