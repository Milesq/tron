package displayer

import (
	tm "github.com/buger/goterm"

	"github.com/milesq/tron/tron"
)

// ConsoleDisplayer .
func ConsoleDisplayer(tr tron.Game) {
	for player, trace := range tr.State.Players {
		for _, pt := range trace {
			tm.MoveCursor(int(pt.X), int(pt.Y))
			tm.Print(tm.Color(string(tr.Cfg.BorderChar), player))
		}

		currentPoint := trace[len(trace)-1]
		tm.MoveCursor(int(currentPoint.X), int(currentPoint.Y))
		tm.Print(tm.Color(string(tr.Cfg.PlayerChar), player))
	}

	tm.MoveCursor(1, 1)
}
