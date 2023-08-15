package solver

import (
	"maze-solver/maze"
	"maze-solver/utils"
	"sort"
)

type AStarSolver struct {
	dist_from_start map[*maze.Node]int
	dist_from_end   map[*maze.Node]int
	parent          map[*maze.Node]*maze.Node
	stack           sorted_stack
}

func (s *AStarSolver) Solve(m *maze.Maze) *maze.SolvedMaze {
	defer utils.Timer("A* algorithm", 2)()
	s.dist_from_start = make(map[*maze.Node]int, len(m.Nodes))
	s.dist_from_end = make(map[*maze.Node]int, len(m.Nodes))
	s.parent = make(map[*maze.Node]*maze.Node, len(m.Nodes))

	current, end := m.Nodes[0], m.Nodes[len(m.Nodes)-1]

	for _, node := range m.Nodes {
		s.dist_from_start[node] = 0
		s.dist_from_end[node] = int(node.Coords.Distance(end.Coords))
	}

	for current != end {
		current.Visited = true

		for _, child := range []*maze.Node{current.Left, current.Right, current.Up, current.Down} {
			if child != nil {
				dist := s.dist_from_start[current] + int(current.Coords.Distance(child.Coords))
				if !child.Visited {
					s.parent[child] = current
					s.dist_from_start[child] = dist
					s.stack.insert(child, &s.dist_from_end)
				} else if s.dist_from_start[child] > dist {
					s.parent[child] = current
					s.dist_from_start[child] = dist
					sort.Slice(s.stack, func(i, j int) bool {
						return s.dist_from_end[s.stack[i]] < s.dist_from_end[s.stack[j]]
					})
				}
			}
		}
		current = s.stack.pop()
	}

	solution := make([]*maze.Node, 0, len(m.Nodes))
	for current != m.Nodes[0] {
		solution = append(solution, current)
		current = s.parent[current]
	}
	solution = append(solution, m.Nodes[0])

	for i, j := 0, len(solution)-1; i < j; i, j = i+1, j-1 {
		solution[i], solution[j] = solution[j], solution[i]
	}

	return &maze.SolvedMaze{
		Maze:     m,
		Solution: solution,
	}
}
