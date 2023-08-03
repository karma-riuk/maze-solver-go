package writer

import (
	"maze-solver/maze"
)

type ImageWriter struct{}

func (w *ImageWriter) Write(filename string, maze *maze.SolvedMaze) error {
	return nil
}
