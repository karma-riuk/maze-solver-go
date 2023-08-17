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

func (f *SolverFactory) Get(solved_chan chan<- *maze.SolvedMaze) Solver {
	switch *f.Type {
	case _DFS:
		return &DFSSolver{
			solved_chan: solved_chan,
		}
	case _BFS:
		return &BFSSolver{
			solved_chan: solved_chan,
		}
	case _AStar:
		return &AStarSolver{
			solved_chan: solved_chan,
		}
	case _Dijkstra:
		return &DijkstraSolver{
			solved_chan: solved_chan,
		}
	}
	panic(fmt.Sprintf("Unrecognized solver type %q", *f.Type))
}

func visited(node *maze.Node) bool {
	return node == nil || node.Visited
}
