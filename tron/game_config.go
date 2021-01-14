package tron

// Config .
type Config struct {
	Players     []string
	PlayerSpeed int
	Size        [2]int
	PlayerChar  byte
	BorderChar  byte
}

// Coords .
type Coords struct {
	X int
	Y int
}

// Trace .
type Trace []Coords

// GameState .
type GameState struct {
	Players map[string][]Trace
	MapSize [2]int
}
