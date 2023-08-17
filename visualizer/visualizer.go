package visualizer

import (
	"image/color"
	"maze-solver/io/writer"
	"maze-solver/maze"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/mazznoer/colorgrad"
)

var (
	a          fyne.App
	w          fyne.Window
	img_writer writer.ImageWriter
	cimg       *canvas.Image
)

func Init(m *maze.Maze) {
	a = app.New()
	w = a.NewWindow("maze-solver-go")
	img_writer = writer.ImageWriter{
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
	println(m.Height)
	cimg = canvas.NewImageFromImage(img_writer.GenerateImage())
	w.SetContent(cimg)
	w.Resize(
		fyne.NewSize(
			float32(m.Width*img_writer.CellWidth),
			float32(m.Height*img_writer.CellHeight),
		),
	)
	w.Show()
}

func Visualize(solved_chan <-chan *maze.SolvedMaze) {
	for solved := range solved_chan {
		img_writer.Maze = solved
		cimg.Image = img_writer.GenerateImage()
		cimg.Refresh()
	}
}

func Run() {
	w.ShowAndRun()
}
