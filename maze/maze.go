package maze

import "math"

type Coordinates struct {
	X, Y int
}

func (c Coordinates) Distance(o Coordinates) float64 {
	x, y := float64(o.X-c.X), float64(o.Y-c.Y)

	if y == 0 {
		if x < 0 {
			return -x
		}
		return x
	}

	if x == 0 {
		if y < 0 {
			return -y
		}
		return y
	}

	return math.Sqrt(x*x + y*y)
}

type Node struct {
	Coords      Coordinates
	Up, Down    *Node
	Left, Right *Node
	Visited     bool `default:"false"`
}

func NewNode(coords Coordinates) *Node {
	return &Node{
		Coords: coords,
		Up:     nil,
		Down:   nil,
		Left:   nil,
		Right:  nil,
	}
}

type Maze struct {
	Width, Height int
	Nodes         []*Node
}

type SolvedMaze struct {
	*Maze
	Solution []*Node
}
