package tron

// PlayerLost .
func (tr *Game) PlayerLost(playerID, conflictedWith int) {
	delete(tr.State.Players, playerID)

	if playerID != conflictedWith {
		tr.State.Points[conflictedWith] += 10
	} else {
		tr.State.Points[playerID] -= 5
	}
}
