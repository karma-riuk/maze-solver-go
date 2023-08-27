package visualizer

import (
	"image/color"
	"maze-solver/io/writer"
	"maze-solver/maze"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"github.com/mazznoer/colorgrad"
)

type WindowVisualizer struct {
	app        fyne.App
	window     fyne.Window
	img_writer writer.ImageWriter
	cimg       *canvas.Image
}

func (v *WindowVisualizer) Init(m *maze.Maze) {
	v.app = app.New()
	v.window = v.app.NewWindow("maze-solver-go")
	v.img_writer = writer.ImageWriter{
		Filename: "",
		Maze: &maze.SolvedMaze{
			Maze:     m,
			Solution: []*maze.Node{},
		},
		CellWidth:        2,
		CellHeight:       2,
		WallColor:        color.Black,
		PathColor:        color.White,
		SolutionGradient: colorgrad.Warm(),
	}
	v.cimg = canvas.NewImageFromImage(v.img_writer.GenerateImage())
	v.window.SetContent(v.cimg)
	v.window.Resize(
		fyne.NewSize(
			m.Width*v.img_writer.CellWidth,
			m.Height*v.img_writer.CellHeight,
		),
	)
	v.window.Show()
}

func (v *WindowVisualizer) Visualize(solved_chan <-chan *maze.SolvedMaze) {
	for solved := range solved_chan {
		v.img_writer.Maze = solved
		v.cimg.Image = v.img_writer.GenerateImage()
		v.cimg.Refresh()
	}
}

func (v *WindowVisualizer) Run(lets_go chan<- bool) {
	dial := dialog.NewConfirm("Start", "Let's go", func(ok bool) {
		lets_go <- ok
		if !ok {
			v.window.Close()
		}
	}, v.window)
	dial.Show()
	v.window.ShowAndRun()
}
