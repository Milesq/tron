package tron

import (
	"math"
)

// Next .
func (tron *Game) Next() {
	for playerID := range tron.State.Players {
		trace := tron.State.Players[playerID]
		movement := tron.PlayersDirection[playerID]
		lastPos := Vector(trace[len(trace)-1])
		newPos := lastPos.Add(movement)

		if roundPt(lastPos) == roundPt(newPos) {
			trace[len(trace)-1] = Point(newPos)
			tron.State.Players[playerID] = trace
		} else { // position has changed
			tron.State.Players[playerID] = append(trace, Point(newPos))

			go func(playerID int) {
				detectConflict(tron.State.Players, newPos, func(id int) {
					tron.PlayerLost(playerID, id)
				})
			}(playerID)
		}
	}
}

func detectConflict(traces map[int]Trace, possibleConflict Vector, conflictHandler func(int)) {
	possibleConflict = roundPt(possibleConflict)

	for playerID, trace := range traces {
		for _, pt := range trace[:len(trace)-1] {
			if roundPt(Vector(pt)) == roundPt(possibleConflict) {
				conflictHandler(playerID)
				return
			}
		}
	}
}

func roundPt(pt Vector) Vector {
	return Vector{math.Floor(pt.X), math.Floor(pt.Y)}
}
