package parser

import (
	"fmt"
	"maze-solver/io/reader"
	"maze-solver/maze"
)

const (
	WallChar = '#'
	PathChar = ' '
)

func Parse(reader reader.Reader) (*maze.Maze, error) {
	nodesByCoord := make(map[maze.Coordinates]*maze.Node)
	ret := &maze.Maze{}

	raw_maze, err := reader.Read()
	if err != nil {
		return nil, err
	}

	for y, line := range raw_maze.Data {
		fmt.Println(line)
		for x := 1; x < len(line)-1; x++ {
			char := line[x]
			var left_char, right_char, above_char byte

			if y > 0 {
				left_char = line[x-1]
				right_char = line[x+1]
				above_char = raw_maze.Data[y-1][x]
			}

			// Parse first line to get entrance
			if y == 0 && char == PathChar {
				coords := maze.Coordinates{X: x, Y: y}
				node := maze.NewNode(coords)
				ret.Nodes = append(ret.Nodes, node)
				nodesByCoord[coords] = node
				continue
			}

			// Parse middle of the maze
			if y > 0 && char == PathChar &&
				(left_char == WallChar && right_char == PathChar ||
					left_char == PathChar && right_char == WallChar ||
					above_char == PathChar && (left_char == PathChar || right_char == PathChar)) {
				coords := maze.Coordinates{X: x, Y: y}
				node := maze.NewNode(coords)

				lookupNeighbourAbove(&raw_maze.Data, node, &nodesByCoord, ret)

				ret.Nodes = append(ret.Nodes, node)
				nodesByCoord[coords] = node

				if left_char == PathChar && right_char == WallChar ||
					above_char == PathChar && (left_char == PathChar || right_char == PathChar) {
					lookupNeighbourLeft(&line, node, &nodesByCoord)
				}
			}
		}
	}

	// Parse last line to get exit
	for x, rune := range raw_maze.Data[len(raw_maze.Data)-1] {
		char := byte(rune)
		if char == PathChar {
			coords := maze.Coordinates{X: x, Y: len(raw_maze.Data) - 1}
			node := maze.NewNode(coords)
			lookupNeighbourAbove(&raw_maze.Data, node, &nodesByCoord, ret)
			ret.Nodes = append(ret.Nodes, node)
			break
		}
	}

	return ret, nil
}

func lookupNeighbourAbove(Data *[]string, node *maze.Node, nodesByCoord *map[maze.Coordinates]*maze.Node, m *maze.Maze) {
	for y := node.Coords.Y - 1; y >= 0; y-- {
		neighbour, ok := (*nodesByCoord)[maze.Coordinates{X: node.Coords.X, Y: y}]

		if ok {
			node.Up = neighbour
			neighbour.Down = node
			break
		}

		if y > 0 && (*Data)[y][node.Coords.X] == WallChar {
			y++
			if y == node.Coords.Y {
				break
			}
			coords := maze.Coordinates{X: node.Coords.X, Y: y}
			new_node := maze.NewNode(coords)
			lookupNeighbourLeft(&(*Data)[y], new_node, nodesByCoord)
			lookupNeighbourRight(&(*Data)[y], new_node, nodesByCoord)
			(*nodesByCoord)[coords] = new_node
			m.Nodes = append(m.Nodes, new_node)

			node.Up = new_node
			new_node.Down = node
			break
		}

	}
}

func lookupNeighbourLeft(line *string, node *maze.Node, nodesByCoord *map[maze.Coordinates]*maze.Node) {
	for x := node.Coords.X - 1; x > 0; x-- {
		if (*line)[x] == WallChar && x < node.Coords.X-1 {
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

func lookupNeighbourRight(line *string, node *maze.Node, nodesByCoord *map[maze.Coordinates]*maze.Node) {
	for x := node.Coords.X + 1; x < len(*line); x++ {
		if (*line)[x] == WallChar {
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
