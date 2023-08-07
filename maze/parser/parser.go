package parser

import (
	"fmt"
	"maze-solver/io/reader"
	"maze-solver/maze"
)

func Parse(reader reader.Reader) (*maze.Maze, error) {
	nodesByCoord := make(map[maze.Coordinates]*maze.Node)
	ret := &maze.Maze{}

	raw_maze, err := reader.Read()
	if err != nil {
		return nil, err
	}

	y := 0
	// Parse first line to get entrance
	for x := 0; x < raw_maze.Width-1; x++ {
		if raw_maze.IsPath(x, y) {
			coords := maze.Coordinates{X: x, Y: y}
			node := maze.NewNode(coords)
			ret.Nodes = append(ret.Nodes, node)
			nodesByCoord[coords] = node
			break
		}
	}

	for y = 1; y < raw_maze.Height-1; y++ {
		for x := 1; x < raw_maze.Width-1; x++ {
			// Parse middle of the maze
			if isCoordEligibleForNode(x, y, raw_maze) {
				coords := maze.Coordinates{X: x, Y: y}
				node := maze.NewNode(coords)

				lookupNeighbourAbove(raw_maze, node, &nodesByCoord, ret)

				ret.Nodes = append(ret.Nodes, node)
				nodesByCoord[coords] = node

				if raw_maze.IsPath(x-1, y) && raw_maze.IsWall(x+1, y) ||
					raw_maze.IsPath(x, y-1) &&
						(raw_maze.IsPath(x-1, y) || raw_maze.IsPath(x+1, y)) {
					lookupNeighbourLeft(raw_maze, node, &nodesByCoord)
				}
			}
		}
	}

	// Parse last line to get exit
	for x := 0; x < raw_maze.Width-1; x++ {
		if raw_maze.IsPath(x, y) {
			coords := maze.Coordinates{X: x, Y: y}
			node := maze.NewNode(coords)
			lookupNeighbourAbove(raw_maze, node, &nodesByCoord, ret)
			ret.Nodes = append(ret.Nodes, node)
			break
		}
	}

	return ret, nil
}

func isCoordEligibleForNode(x int, y int, raw_maze *reader.RawMaze) bool {
	return raw_maze.IsPath(x, y) &&
		(raw_maze.IsWall(x-1, y) && raw_maze.IsPath(x+1, y) || // wall left, path right
			raw_maze.IsPath(x-1, y) && raw_maze.IsWall(x+1, y) || // path left, wall right
			raw_maze.IsPath(x, y-1) && (raw_maze.IsPath(x-1, y) || raw_maze.IsPath(x+1, y)) || // path above and not in vertical corridor
			raw_maze.IsWall(x-1, y) && raw_maze.IsWall(x+1, y) && raw_maze.IsPath(x, y-1) && raw_maze.IsWall(x, y+1)) // wall to left, below, above and path above
}

func lookupNeighbourAbove(raw_maze *reader.RawMaze, node *maze.Node, nodesByCoord *map[maze.Coordinates]*maze.Node, m *maze.Maze) {
	for y := node.Coords.Y - 1; y >= 0; y-- {
		neighbour, ok := (*nodesByCoord)[maze.Coordinates{X: node.Coords.X, Y: y}]

		if ok {
			node.Up = neighbour
			neighbour.Down = node
			break
		}

		if y > 0 && raw_maze.IsWall(node.Coords.X, y) {
			y++
			if y == node.Coords.Y {
				break
			}
			coords := maze.Coordinates{X: node.Coords.X, Y: y}
			new_node := maze.NewNode(coords)
			lookupNeighbourLeft(raw_maze, new_node, nodesByCoord)
			lookupNeighbourRight(raw_maze, new_node, nodesByCoord)
			(*nodesByCoord)[coords] = new_node
			m.Nodes = append(m.Nodes, new_node)

			node.Up = new_node
			new_node.Down = node
			break
		}

	}
}

func lookupNeighbourLeft(raw_maze *reader.RawMaze, node *maze.Node, nodesByCoord *map[maze.Coordinates]*maze.Node) {
	for x := node.Coords.X - 1; x > 0; x-- {
		if raw_maze.IsWall(x, node.Coords.Y) && x < node.Coords.X-1 {
			panic(fmt.Sprintf("Found no node before wall while looking to the left at neighbours of node %v (arrived at x=%v before hitting a wall)", node, x))
		}

		neighbour, ok := (*nodesByCoord)[maze.Coordinates{X: x, Y: node.Coords.Y}]
		if ok {
			node.Left = neighbour
			neighbour.Right = node
			break
		}
	}
}

func lookupNeighbourRight(raw_maze *reader.RawMaze, node *maze.Node, nodesByCoord *map[maze.Coordinates]*maze.Node) {
	for x := node.Coords.X + 1; x < raw_maze.Width; x++ {
		if raw_maze.IsWall(x, node.Coords.Y) {
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
