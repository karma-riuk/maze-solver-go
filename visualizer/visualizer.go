package visualizer

import (
	"fmt"
	"maze-solver/maze"
)

type Visualizer interface {
	Init(*maze.Maze)
	Visualize(<-chan *maze.SolvedMaze)
	Run(lets_go chan<- bool)
}

type VisualizerFactory struct {
	Type      *string
	Filename  *string
	Framerate *float64
}

const (
	_VIDEO  = "video"
	_WINDOW = "window"
)

var VIZ_METHODS = []string{
	_VIDEO,
	_WINDOW,
}

func (f *VisualizerFactory) Get() Visualizer {
	switch *f.Type {
	case _VIDEO:
		return &VideoVisualizer{
			Filename:  *f.Filename,
			Framerate: *f.Framerate,
		}
	case _WINDOW:
		return &WindowVisualizer{}
	}
	panic(fmt.Sprintf("Unrecognized visualizer type %q", *f.Type))
}
