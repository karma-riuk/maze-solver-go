package reader

import "maze-solver/maze"

type Reader interface {
	Read(filename string) (*maze.Maze, error)
}
