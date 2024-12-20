package commons

import (
	"fmt"
	"math"
)

type Coord struct {
	X int
	Y int
}

func (c Coord) Add(other Coord) Coord {
	return Coord{c.X + other.X, c.Y + other.Y}
}

func (c Coord) Sub(other Coord) Coord {
	return Coord{c.X - other.X, c.Y - other.Y}
}

func (c Coord) Scale(factor int) Coord {
	return Coord{c.X * factor, c.Y * factor}
}

func (c Coord) WrapAround(xSize, ySize int) Coord {
	x := c.X
	y := c.Y
	// Modulo doesn't work correctly with negative numbers
	if x < 0 {
		x += int(math.Ceil(math.Abs(float64(x)/float64(xSize))))*xSize
	}
	if y < 0 {
		y += int(math.Ceil(math.Abs(float64(y)/float64(ySize))))*ySize
	}

	return Coord{x % xSize, y % ySize}
}

func (c Coord) Move(d Direction) Coord {
	return c.Add(Offsets[d])
}

func (c1 Coord) Equals(c2 Coord) bool {
	return c1.X == c2.X && c1.Y == c2.Y
}

func (c Coord) Inbounds(xSize, ySize int) bool {
	return c.X >= 0 && c.Y >= 0 && c.X <= xSize && c.Y <= ySize
}

// Uses first index as y and second as x
func IndexMap[T any](m [][]T, c Coord) T {
	return m[c.Y][c.X]
}

func SetMap[T any](m *[][]T, c Coord, t T) {
	(*m)[c.Y][c.X] = t
}

func (c Coord) String() string {
	return "(" + fmt.Sprint(c.X) + ", " + fmt.Sprint(c.Y) + ")"
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

var Directions = []Direction{
	UP,
	DOWN,
	LEFT,
	RIGHT,
}

// Suposes (0,0) at top-left corner, with y going downwards
// and x going to the right
var Offsets = map[Direction]Coord{
	UP:    {Y: -1, X: 0},
	DOWN:  {Y: 1, X: 0},
	LEFT:  {Y: 0, X: -1},
	RIGHT: {Y: 0, X: 1},
}

var DirFromOffset = map[Coord]Direction{
	{Y: -1, X: 0}: UP,
	{Y: 1, X: 0}:  DOWN,
	{Y: 0, X: -1}: LEFT,
	{Y: 0, X: 1}:  RIGHT,
}

var InverseDir = map[Direction]Direction{
	UP:    DOWN,
	DOWN:  UP,
	LEFT:  RIGHT,
	RIGHT: LEFT,
}

var RotateClockwise = map[Direction]Direction{
	UP:    RIGHT,
	RIGHT: DOWN,
	DOWN:  LEFT,
	LEFT:  UP,
}

var RotateCounterClockwise = map[Direction]Direction{
	UP:    LEFT,
	RIGHT: UP,
	DOWN:  RIGHT,
	LEFT:  DOWN,
}

var OrthogonalDirections = map[Direction][]Direction{
	UP:    {LEFT, RIGHT},
	DOWN:  {LEFT, RIGHT},
	LEFT:  {UP, DOWN},
	RIGHT: {UP, DOWN},
}

func (d Direction) String() string {
	switch d {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	}
	panic("Unreachable")
}
