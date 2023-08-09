package writer

import (
	"bytes"
	"fmt"
	"maze-solver/maze"
)

type StringsWriter struct {
	PathChar, WallChar byte
	SolutionChar       byte
	Maze               *maze.SolvedMaze
	lines              [][]byte
}

func (w *StringsWriter) Write() error {
	w.lines = make([][]byte, w.Maze.Height)
	for y := 0; y < w.Maze.Height; y++ {
		w.lines[y] = bytes.Repeat([]byte{w.WallChar}, w.Maze.Width)
	}
	for _, node := range w.Maze.Nodes {
		if node.Right != nil {
			w.fillHorizontally(node.Coords, node.Right.Coords, w.PathChar)
		}
		if node.Down != nil {
			w.fillVertically(node.Coords, node.Down.Coords, w.PathChar)
		}
	}

	for i := 0; i < len(w.Maze.Solution)-1; i++ {
		current := w.Maze.Solution[i].Coords
		next := w.Maze.Solution[i+1].Coords

		if current.X == next.X {
			w.fillVertically(current, next, w.SolutionChar)
		} else {
			w.fillHorizontally(current, next, w.SolutionChar)
		}
	}

	return nil
}

func (w *StringsWriter) fillHorizontally(from maze.Coordinates, to maze.Coordinates, char byte) {
	y := from.Y
	if from.X > to.X {
		from, to = to, from
	}
	for x := from.X; x <= to.X; x++ {
		w.lines[y][x] = char
	}
}

func (w *StringsWriter) fillVertically(from maze.Coordinates, to maze.Coordinates, char byte) {
	x := from.X
	for y := from.Y; y <= to.Y; y++ {
		w.lines[y][x] = char
	}
}

func (w *StringsWriter) GetLines() []string {
	ret := make([]string, len(w.lines))
	for i, line := range w.lines {
		ret[i] = fmt.Sprint(string(line))
	}
	return ret
}
