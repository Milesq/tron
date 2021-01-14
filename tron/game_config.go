package tron

// Config .
type Config struct {
	Players     []string
	PlayerSpeed int
	Size        [2]int
	PlayerChar  byte
	BorderChar  byte
}

// Point .
type Point struct {
	X int
	Y int
}

// Vector .
type Vector Point

// Trace .
type Trace []Point

// GameState .
type GameState struct {
	Players map[string]Trace
	MapSize [2]int
}
