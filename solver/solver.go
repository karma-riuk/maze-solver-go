package solver

import "maze-solver/maze"

type Solver interface {
	Solve(*maze.Maze) *maze.SolvedMaze
}
