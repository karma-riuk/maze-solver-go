package writer

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"maze-solver/maze"
	"maze-solver/utils"
	"os"

	"github.com/mazznoer/colorgrad"
)

type ImageWriter struct {
	Filename              string
	Maze                  *maze.SolvedMaze
	CellWidth, CellHeight int
	WallColor, PathColor  color.Color
	SolutionGradient      colorgrad.Gradient
	img                   *image.RGBA
}

func (w *ImageWriter) Write() error {
	defer utils.Timer("Image writer", 3)()
	if w.Filename[len(w.Filename)-4:] != ".png" {
		return errors.New("Filename of ImageWriter doesn't have .png extension. The only suppported image type is png")
	}

	w.GenerateImage()

	f, err := os.Create(w.Filename)
	if err != nil {
		return err
	}
	png.Encode(f, w.img)
	f.Close()
	return nil
}

func (w *ImageWriter) GenerateImage() *image.RGBA {
	w.img = image.NewRGBA(image.Rect(0, 0, w.Maze.Width*w.CellWidth, w.Maze.Height*w.CellHeight))

	// Fill the image with walls
	draw.Draw(w.img, w.img.Bounds(), &image.Uniform{w.WallColor}, image.Pt(0, 0), draw.Src)

	// Fill in the paths
	var x0, y0, width, height int
	for _, node := range w.Maze.Nodes {
		x0 = node.Coords.X * w.CellWidth
		y0 = node.Coords.Y * w.CellHeight
		if node.Right != nil {
			width = (node.Right.Coords.X - node.Coords.X + 1) * w.CellWidth
			height = w.CellHeight
			w.draw(x0, y0, width, height, w.PathColor)
		}

		if node.Down != nil {
			width = w.CellWidth
			height = (node.Down.Coords.Y - node.Coords.Y + 1) * w.CellHeight
			w.draw(x0, y0, width, height, w.PathColor)
		}
	}
	if len(w.Maze.Solution) == 0 {
		return w.img
	}

	// Fill in the solution
	total_len := w.getSolutionLength()
	colors := w.SolutionGradient.Colors(uint(total_len + 1))
	c := 0
	width, height = w.CellWidth, w.CellHeight
	for i, from := range w.Maze.Solution[:len(w.Maze.Solution)-1] {
		to := w.Maze.Solution[i+1]

		if from.Coords.X == to.Coords.X {
			// Fill verticallly
			x0 = from.Coords.X * w.CellWidth

			if from.Coords.Y < to.Coords.Y {
				for y := from.Coords.Y; y < to.Coords.Y; y++ {
					y0 = y * w.CellHeight
					w.draw(x0, y0, width, height, colors[c])
					c++
				}
			} else {
				for y := from.Coords.Y; y > to.Coords.Y; y-- {
					y0 = y * w.CellHeight
					w.draw(x0, y0, width, height, colors[c])
					c++
				}
			}
			y0 = to.Coords.Y * w.CellHeight
			w.draw(x0, y0, width, height, colors[c])
		} else {
			// Fill horizontally
			y0 = from.Coords.Y * w.CellHeight

			if from.Coords.X < to.Coords.X {
				for x := from.Coords.X; x < to.Coords.X; x++ {
					x0 = x * w.CellWidth
					w.draw(x0, y0, width, height, colors[c])
					c++
				}
			} else {
				for x := from.Coords.X; x > to.Coords.X; x-- {
					x0 = x * w.CellWidth
					w.draw(x0, y0, width, height, colors[c])
					c++
				}
			}
			x0 = to.Coords.X * w.CellWidth
			w.draw(x0, y0, width, height, colors[c])
		}
	}
	return w.img
}

func (w *ImageWriter) getSolutionLength() int {
	ret := 0
	for i, node := range w.Maze.Solution[:len(w.Maze.Solution)-1] {
		next := w.Maze.Solution[i+1]
		ret += int(node.Coords.Distance(next.Coords))
	}
	return ret
}

func (w *ImageWriter) draw(x0, y0, width, height int, color color.Color) {
	draw.Draw(w.img, image.Rect(x0, y0, x0+width, y0+height), &image.Uniform{color}, image.Pt(0, 0), draw.Src)
}
