package tron

// Config .
type Config struct {
	Players     []int // []tm.Color
	PlayerSpeed float64
	Size        [2]int
	PlayerChar  byte
	BorderChar  byte
}

// Point .
type Point struct {
	X float64
	Y float64
}

// Vector .
type Vector Point

var (
	// UP .
	UP = Vector{0, -1}

	// RIGHT .
	RIGHT = Vector{1, 0}

	// DOWN .
	DOWN = Vector{0, 1}

	// LEFT .
	LEFT = Vector{-1, 0}
)

// Mul miltiplies two vectors
func (v Vector) Mul(second Vector) Vector {
	return Vector{
		v.X * second.X,
		v.Y * second.Y,
	}
}

// Trace .
type Trace []Point

// GameState .
type GameState struct {
	Players map[int]Trace
}
