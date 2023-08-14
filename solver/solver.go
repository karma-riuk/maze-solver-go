package solver

import (
	"fmt"
	"maze-solver/maze"
)

type Solver interface {
	Solve(*maze.Maze) *maze.SolvedMaze
}

type solver struct {
	visited map[*maze.Node]bool
}

type SolverFactory struct {
	Type *string
}

const (
	_DFS = "dfs"
	_BFS = "bfs"
)

var TYPES = []string{
	_DFS,
	_BFS,
}

func (f *SolverFactory) Get() Solver {
	switch *f.Type {
	case _DFS:
		return &DFSSolver{}
	case _BFS:
		return &BFSSolver{}
	}
	panic(fmt.Sprintf("Unrecognized solver type %q", *f.Type))
}

func (s *solver) wasVisited(node *maze.Node) bool {
	if node == nil {
		return true
	}
	visited, _ := s.visited[node]
	return visited
}

func (s *solver) initVisited(m *maze.Maze) {
	s.visited = make(map[*maze.Node]bool, len(m.Nodes))

	for _, node := range m.Nodes {
		s.visited[node] = false
	}
}
