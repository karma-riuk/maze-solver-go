package reader

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

type ImageReader struct {
	Filename              string
	PathColor, WallColor  color.Color
	CellWidth, CellHeight int
}

func (r *ImageReader) Read() (*RawMaze, error) {
	image, err := r.getShrunkImage()
	if err != nil {
		return nil, err
	}

	width, height := image.Bounds().Max.X, image.Bounds().Max.Y
	ret := &RawMaze{
		Width:  width,
		Height: height,
		Data:   make([][]byte, height),
	}

	n_chunks := width/CHUNK_SIZE + 1

	for i := 0; i < height; i++ {
		ret.Data[i] = make([]byte, n_chunks)
	}

	for y := 0; y < height; y++ {
		for i := 0; i < n_chunks; i++ {
			var chunk byte = 0 // all walls

			end_index := min((i+1)*CHUNK_SIZE, width)

			for x := i * CHUNK_SIZE; x < end_index; x++ {
				c := image.At(x, y)
				if c == r.PathColor {
					chunk |= 1 << (CHUNK_SIZE - 1 - (x - i*CHUNK_SIZE))
				}
			}

			ret.Data[y][i] = chunk
		}
	}

	return ret, nil
}

func (r *ImageReader) getShrunkImage() (*image.RGBA, error) {
	input, err := os.Open(r.Filename)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	// Decode the image (from PNG to image.Image):
	src, _ := png.Decode(input)

	// Set the expected size that you want:
	dst := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/r.CellWidth, src.Bounds().Max.Y/r.CellHeight))

	// Resize:
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	return dst, nil
}
