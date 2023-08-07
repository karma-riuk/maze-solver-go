package maze

type Coordinates struct {
	X, Y int
}
type Node struct {
	Coords      Coordinates
	Up, Down    *Node
	Left, Right *Node
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
	Maze
	Solution []*Node
}
