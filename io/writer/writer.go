package writer

import (
	"fmt"
	"image/color"
	"maze-solver/maze"

	"github.com/mazznoer/colorgrad"
)

type Writer interface {
	Write() error
}

type WriterFactory struct {
	Type                             string
	Filename                         *string
	PathChar, WallChar, SolutionChar *string
	CellWidth, CellHeight            *int
	WallColor, PathColor             color.Color
	SolutionGradient                 colorgrad.Gradient
}

const (
	_IMAGE = "image"
)

var TYPES = map[string]string{
	".png": _IMAGE,
}

func (f *WriterFactory) Get(m *maze.SolvedMaze) Writer {
	switch f.Type {
	case _IMAGE:
		return &ImageWriter{
			Filename:         *f.Filename,
			Maze:             m,
			CellWidth:        *f.CellWidth,
			CellHeight:       *f.CellHeight,
			WallColor:        f.WallColor,
			PathColor:        f.PathColor,
			SolutionGradient: f.SolutionGradient,
		}
	}
	panic(fmt.Sprintf("Unrecognized writer type %q", f.Type))
}
