package reader

import (
	"fmt"
	"image/color"
)

type Reader interface {
	Read() (*RawMaze, error)
}

type ReaderFactory struct {
	Type                             string
	Filename                         *string
	PathChar, WallChar, SolutionChar *string
	CellWidth, CellHeight            *int
	WallColor, PathColor             color.Color
}

const (
	_IMAGE = "image"
	_TEXT  = "text"
)

var TYPES = map[string]string{
	".png": _IMAGE,
	".txt": _TEXT,
}

func (f *ReaderFactory) Get() Reader {
	switch f.Type {
	case _TEXT:
		return &TextReader{
			Filename: *f.Filename,
			PathChar: byte((*f.PathChar)[0]),
			WallChar: byte((*f.WallChar)[0]),
		}
	case _IMAGE:
		return &ImageReader{
			Filename:   *f.Filename,
			CellWidth:  *f.CellWidth,
			CellHeight: *f.CellHeight,
			WallColor:  f.WallColor,
			PathColor:  f.PathColor,
		}
	}
	panic(fmt.Sprintf("Unrecognized reader type %q", f.Type))
}
