package writer

import "maze-solver/maze"

type Writer interface {
	Write(filename string, maze *maze.SolvedMaze) error
}
