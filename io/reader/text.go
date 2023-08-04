package reader

import (
	"bufio"
	"fmt"
	"maze-solver/maze"
	"os"
)

type TextReader struct {
	PathChar, WallChar byte
}

func (r *TextReader) Read(filename string) (*maze.Maze, error) {
	nodesByCoord := make(map[maze.Coordinates]*maze.Node)
	var lines []string

	ret := &maze.Maze{}

	if _, err := os.Stat(filename); err != nil {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	y := 0
	var line string
	for scanner.Scan() {
		line = scanner.Text()

		if len(lines) == 0 {
			lines = make([]string, 0, len(line))
		}

		for x := 1; x < len(line)-1; x++ {
			char := line[x]
			var left_char, right_char, above_char byte

			if y > 0 {
				left_char = line[x-1]
				right_char = line[x+1]
				above_char = lines[y-1][x]
			}

			// Parse first line to get entrance
			if y == 0 && char == r.PathChar {
				coords := maze.Coordinates{X: x, Y: y}
				node := maze.NewNode(coords)
				ret.Nodes = append(ret.Nodes, node)
				nodesByCoord[coords] = node
				continue
			}

			// Parse middle of the maze
			if y > 0 && char == r.PathChar &&
				(left_char == r.WallChar && right_char == r.PathChar ||
					left_char == r.PathChar && right_char == r.WallChar ||
					above_char == r.PathChar && (left_char == r.PathChar || right_char == r.PathChar)) {
				coords := maze.Coordinates{X: x, Y: y}
				node := maze.NewNode(coords)
				ret.Nodes = append(ret.Nodes, node)
				nodesByCoord[coords] = node

				r.lookupNeighbourAbove(&lines, node, &nodesByCoord, ret)
				if left_char == r.PathChar && right_char == r.WallChar ||
					above_char == r.PathChar && (left_char == r.PathChar || right_char == r.PathChar) {
					r.lookupNeighbourLeft(&line, node, &nodesByCoord)
				}
			}
		}
		lines = append(lines, line)
		y++
	}
	y--
	// Parse last line to get exit
	for x, rune := range line {
		char := byte(rune)
		if char == r.PathChar {
			coords := maze.Coordinates{X: x, Y: y}
			node := maze.NewNode(coords)
			r.lookupNeighbourAbove(&lines, node, &nodesByCoord, ret)
			ret.Nodes = append(ret.Nodes, node)
			break
		}
	}

	return ret, nil
}

func (r *TextReader) lookupNeighbourAbove(lines *[]string, node *maze.Node, nodesByCoord *map[maze.Coordinates]*maze.Node, m *maze.Maze) {
	for y := node.Coords.Y - 1; y >= 0; y-- {
		neighbour, ok := (*nodesByCoord)[maze.Coordinates{X: node.Coords.X, Y: y}]

		if ok {
			node.Up = neighbour
			neighbour.Down = node
			break
		}

		if y > 0 && (*lines)[y][node.Coords.X] == r.WallChar {
			y++
			if y == node.Coords.Y {
				break
			}
			coords := maze.Coordinates{X: node.Coords.X, Y: y}
			new_node := maze.NewNode(coords)
			r.lookupNeighbourLeft(&(*lines)[y], new_node, nodesByCoord)
			r.lookupNeighbourRight(&(*lines)[y], new_node, nodesByCoord)
			(*nodesByCoord)[coords] = new_node
			m.Nodes = append(m.Nodes, new_node)

			node.Up = new_node
			new_node.Down = node
			break
		}

	}
}

func (r *TextReader) lookupNeighbourLeft(line *string, node *maze.Node, nodesByCoord *map[maze.Coordinates]*maze.Node) {
	for x := node.Coords.X - 1; x > 0; x-- {
		if (*line)[x] == r.WallChar {
			panic(fmt.Sprintf("Found no node before wall while looking to the left at neighbours of node %v", node))
		}

		neighbour, ok := (*nodesByCoord)[maze.Coordinates{X: x, Y: node.Coords.Y}]
		if ok {
			node.Left = neighbour
			neighbour.Right = node
			break
		}
	}
}

func (r *TextReader) lookupNeighbourRight(line *string, node *maze.Node, nodesByCoord *map[maze.Coordinates]*maze.Node) {
	for x := node.Coords.X + 1; x < len(*line); x++ {
		if (*line)[x] == r.WallChar {
			panic(fmt.Sprintf("Found no node before wall while looking to the right at neighbours of node %v", node))
		}

		neighbour, ok := (*nodesByCoord)[maze.Coordinates{X: x, Y: node.Coords.Y}]
		if ok {
			node.Right = neighbour
			neighbour.Left = node
			break
		}
	}
}
