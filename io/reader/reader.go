package reader

import "maze-solver/maze"

type Reader interface {
	Read() (*maze.RawMaze, error)
}
