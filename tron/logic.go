package tron

// Next .
func (tron *Game) Next() {
	for playerID := range tron.State.Players {
		trace := tron.State.Players[playerID]
		vector := tron.PlayersDirection[playerID]
		lastPos := Vector(trace[len(trace)-1])

		tron.State.Players[playerID] = append(trace, Point(lastPos.Add(vector)))
	}
}
