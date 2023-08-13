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
	_TURN_LEFT = "turn-left"
)

var TYPES = []string{
	_TURN_LEFT,
}

func (f *SolverFactory) Get() Solver {
	switch *f.Type {
	case _TURN_LEFT:
		return &DFSSolver{}
	}
	panic(fmt.Sprintf("Unrecognized solver type %q", *f.Type))
}
