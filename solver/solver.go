package solver

import (
	"fmt"
	"maze-solver/maze"
)

type Solver interface {
	Solve(*maze.Maze) *maze.SolvedMaze
}

type SolverFactory struct {
	Type *string
}

const (
	_DFS      = "dfs"
	_BFS      = "bfs"
	_Dijkstra = "dijkstra"
	_AStar    = "a-star"
)

var TYPES = []string{
	_DFS,
	_BFS,
	_Dijkstra,
	_AStar,
}

func (f *SolverFactory) Get() Solver {
	switch *f.Type {
	case _DFS:
		return &DFSSolver{}
	case _BFS:
		return &BFSSolver{}
	case _AStar:
		return &AStarSolver{}
	case _Dijkstra:
		return &DijkstraSolver{}
	}
	panic(fmt.Sprintf("Unrecognized solver type %q", *f.Type))
}

func visited(node *maze.Node) bool {
	return node == nil || node.Visited
}
