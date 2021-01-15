package tron

// Config .
type Config struct {
	Players     []int
	PlayerSpeed int // tm.Color
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
	Players map[int]Trace
	MapSize [2]int
}
