package solver

import (
	"maze-solver/maze"
	"maze-solver/utils"
)

type TurnLeftSolver struct {
	visited map[*maze.Node]bool
}

func (s *TurnLeftSolver) Solve(m *maze.Maze) *maze.SolvedMaze {
	defer utils.Timer("Turn left algorithm", 2)()

	current, end := m.Nodes[0], m.Nodes[len(m.Nodes)-1]
	s.visited = make(map[*maze.Node]bool, len(m.Nodes))

	for _, node := range m.Nodes {
		s.visited[node] = false
	}

	stack := make([]*maze.Node, 0, len(m.Nodes))
	stack = append(stack, current)

	for current != end {
		s.visited[current] = true

		left_visited, right_visited, up_visited, down_visited := s.wasVisited(current.Left), s.wasVisited(current.Right), s.wasVisited(current.Up), s.wasVisited(current.Down)

		if left_visited && right_visited && up_visited && down_visited {
			// dead end or no more visited nodes
			stack = stack[:len(stack)-1]
			current = stack[len(stack)-1]
		} else {

			if !left_visited {
				current = current.Left
			} else if !down_visited {
				current = current.Down
			} else if !right_visited {
				current = current.Right
			} else if !up_visited {
				current = current.Up
			}

			stack = append(stack, current)
		}
	}

	ret := &maze.SolvedMaze{
		Maze:     m,
		Solution: stack,
	}
	return ret
}

func (s *TurnLeftSolver) wasVisited(node *maze.Node) bool {
	if node == nil {
		return true
	}
	visited, _ := s.visited[node]
	return visited
}
