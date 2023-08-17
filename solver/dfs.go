package solver

import (
	"maze-solver/maze"
	"maze-solver/utils"
)

type DFSSolver struct {
	solved_chan chan<- *maze.SolvedMaze
}

func (s *DFSSolver) Solve(m *maze.Maze) *maze.SolvedMaze {
	defer utils.Timer("DFS algorithm", 2)()

	current, end := m.Nodes[0], m.Nodes[len(m.Nodes)-1]

	stack := make([]*maze.Node, 0, len(m.Nodes))
	stack = append(stack, current)

	for current != end {
		current.Visited = true
		if s.solved_chan != nil {
			s.solved_chan <- &maze.SolvedMaze{
				Maze:     m,
				Solution: stack,
			}
		}

		left_visited, right_visited, up_visited, down_visited := visited(current.Left), visited(current.Right), visited(current.Up), visited(current.Down)

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
