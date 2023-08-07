package reader

import (
	"maze-solver/utils"
	"reflect"
	"testing"
)

func TestTextReadTrivial(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		pathChar byte
		wallChar byte
		expected *RawMaze
	}{
		{
			"Trivial",
			"../../assets/trivial.txt",
			' ',
			'#',
			&RawMaze{
				Width:  5,
				Height: 3,
				Data: [][]byte{
					{0b_00100_000},
					{0b_01110_000},
					{0b_00010_000},
				},
			},
		},
		{
			"Trivial Bigger",
			"../../assets/trivial-bigger.txt",
			' ',
			'#',
			&RawMaze{
				Width:  7,
				Height: 5,
				Data: [][]byte{
					{0b_0001000_0},
					{0b_0001000_0},
					{0b_0111110_0},
					{0b_0000010_0},
					{0b_0000010_0},
				},
			},
		},
		{
			"Bigger Staggered",
			"../../assets/trivial-bigger-staggered.txt",
			' ',
			'#',
			&RawMaze{
				Width:  7,
				Height: 5,
				Data: [][]byte{
					{0b_0001000_0},
					{0b_0001000_0},
					{0b_0111110_0},
					{0b_0000100_0},
					{0b_0000100_0},
				},
			},
		},
		{
			"Normal",
			"../../assets/normal.txt",
			' ',
			'#',
			&RawMaze{
				Width:  11,
				Height: 11,
				Data: [][]byte{
					{0b_00000100, 0b000_00000},
					{0b_01111101, 0b110_00000},
					{0b_00000100, 0b010_00000},
					{0b_01110111, 0b110_00000},
					{0b_01010000, 0b010_00000},
					{0b_01011111, 0b110_00000},
					{0b_00010001, 0b010_00000},
					{0b_01110111, 0b010_00000},
					{0b_01000000, 0b010_00000},
					{0b_01111101, 0b110_00000},
					{0b_00000100, 0b000_00000},
				},
			},
		},
	}

	for _, test := range tests {
		reader := TextReader{
			Filename: test.filename,
			PathChar: test.pathChar,
			WallChar: test.wallChar,
		}

		got, err := reader.Read()
		utils.Check(err, "Couldn't read file %q", reader.Filename)

		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("%s: lexed mazes do not match\nGot: %v\nWant: %v", test.name, got, test.expected)
		}
	}
}
