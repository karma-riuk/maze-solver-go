package writer

import (
	"bytes"
	"image/color"
	"io"
	"maze-solver/maze"
	"os"
	"testing"

	"github.com/mazznoer/colorgrad"
)

const (
	OUT_DIR      = "./out"
	EXPECTED_DIR = "../../assets/solved"
)

func TestImageWriter(t *testing.T) {
	if _, err := os.Stat(OUT_DIR); os.IsNotExist(err) {
		os.Mkdir(OUT_DIR, 0700)
	}

	tests := []struct {
		name                  string
		filename              string
		m                     *maze.SolvedMaze
		CellWidth, cellHeight int
		pathColor, wallColor  color.Color
		gradient              colorgrad.Gradient
	}{
		{
			"Trivial",
			"trivial.png",
			trivial(),
			20, 20,
			color.White, color.Black,
			colorgrad.Warm(),
		},
		{
			"Trivial Bigger",
			"trivial-bigger.png",
			bigger(),
			20, 20,
			color.White, color.Black,
			colorgrad.Warm(),
		},
		{
			"Trivial Bigger Staggered",
			"trivial-bigger-staggered.png",
			bigger_staggered(),
			20, 20,
			color.White, color.Black,
			colorgrad.Warm(),
		},
		{
			"Normal",
			"normal.png",
			normal(),
			20, 20,
			color.White, color.Black,
			colorgrad.Warm(),
		},
	}

	for _, test := range tests {
		writer := ImageWriter{
			Filename:         OUT_DIR + "/" + test.filename,
			Maze:             test.m,
			CellWidth:        test.CellWidth,
			CellHeight:       test.cellHeight,
			WallColor:        test.wallColor,
			PathColor:        test.pathColor,
			SolutionGradient: test.gradient,
		}

		err := writer.Write()
		if err != nil {
			t.Fatalf("%s: couldn't write solution, got following error\n%v", test.name, err)
		}

		assertEqualFile(t, EXPECTED_DIR+"/"+test.filename, OUT_DIR+"/"+test.filename, test.name)
		os.Remove(writer.Filename)
	}
	os.Remove(OUT_DIR)
}

const chunkSize = 64000

func assertEqualFile(t *testing.T, file1, file2, name string) {
	f1, err := os.Open(file1)
	if err != nil {
		t.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		t.Fatal(err)
	}
	defer f2.Close()

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return
			} else if err1 == io.EOF || err2 == io.EOF {
				t.Fatalf("%s: files are not equal. Got %q, wanted %q", name, file1, file2)
			}
		}

		if !bytes.Equal(b1, b2) {
			t.Fatalf("%s: files are not equal. Got %q, wanted %q", name, file1, file2)
		}
	}
}
