package tron

import "fmt"

// PlayerLost .
func (tr *Game) PlayerLost(playerID, conflictedWith int) {
	delete(tr.State.Players, playerID)
	fmt.Println(tr.State.Players)
	tr.State.Points[conflictedWith]++
}
