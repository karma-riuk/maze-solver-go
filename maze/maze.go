package maze

import (
	"fmt"
	"strings"
)

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

type RawMaze struct {
	PathChar, WallChar byte
	Data               []string
}

func (m *RawMaze) String() string {
	var ret strings.Builder
	ret.WriteString("{\n")
	ret.WriteString(fmt.Sprintf("\tPathChar: %v,\n", m.PathChar))
	ret.WriteString(fmt.Sprintf("\tWallChar: %v,\n", m.WallChar))
	ret.WriteString("\tData: \n")
	for _, line := range m.Data {
		ret.WriteRune('\t')
		ret.WriteRune('\t')
		ret.WriteString(line)
		ret.WriteRune('\n')
	}
	ret.WriteString("}")

	return ret.String()
}

type Maze struct {
	Width, Height uint
	Nodes         []*Node
}

type SolvedMaze struct {
	Maze
	Solution []*Node
}
