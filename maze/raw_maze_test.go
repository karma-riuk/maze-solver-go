package maze

import "testing"

func TestRawMazeWall(t *testing.T) {
	tests := []struct {
		name          string
		width, height int
		data          [][]byte
		expected      [][]bool
	}{
		{
			"Trivial",
			5, 3,
			[][]byte{
				{0b_00100_000},
				{0b_01110_000},
				{0b_00010_000},
			},
			[][]bool{
				{true, true, false, true, true},
				{true, false, false, false, true},
				{true, true, true, false, true},
			},
		},
		{
			"Trivial Bigger",
			7, 5,
			[][]byte{
				{0b_0001000_0},
				{0b_0001000_0},
				{0b_0111110_0},
				{0b_0000010_0},
				{0b_0000010_0},
			},
			[][]bool{
				{true, true, true, false, true, true, true},
				{true, true, true, false, true, true, true},
				{true, false, false, false, false, false, true},
				{true, true, true, true, true, false, true},
				{true, true, true, true, true, false, true},
			},
		},
		{
			"Bigger Staggered",
			7, 5,
			[][]byte{
				{0b_0001000_0},
				{0b_0001000_0},
				{0b_0111110_0},
				{0b_0000100_0},
				{0b_0000100_0},
			},
			[][]bool{
				{true, true, true, false, true, true, true},
				{true, true, true, false, true, true, true},
				{true, false, false, false, false, false, true},
				{true, true, true, true, false, true, true},
				{true, true, true, true, false, true, true},
			},
		},
		{
			"Normal",
			11, 11,
			[][]byte{
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
			[][]bool{
				{true, true, true, true, true, false, true, true, true, true, true},
				{true, false, false, false, false, false, true, false, false, false, true},
				{true, true, true, true, true, false, true, true, true, false, true},
				{true, false, false, false, true, false, false, false, false, false, true},
				{true, false, true, false, true, true, true, true, true, false, true},
				{true, false, true, false, false, false, false, false, false, false, true},
				{true, true, true, false, true, true, true, false, true, false, true},
				{true, false, false, false, true, false, false, false, true, false, true},
				{true, false, true, true, true, true, true, true, true, false, true},
				{true, false, false, false, false, false, true, false, false, false, true},
				{true, true, true, true, true, false, true, true, true, true, true},
			},
		},
	}

	for _, test := range tests {
		rawMaze := RawMaze{
			Width:  test.width,
			Height: test.height,
			Data:   test.data,
		}
		for y, row := range test.expected {
			for x, expected := range row {
				if rawMaze.IsWall(x, y) != expected {
					t.Fatalf("%s: Wanted wall at (%v, %v), apparently it isn't", test.name, x, y)
				}
			}
		}
	}
}

func TestRawMazePath(t *testing.T) {
	tests := []struct {
		name          string
		width, height int
		data          [][]byte
		expected      [][]bool
	}{
		{
			"Trivial",
			5, 3,
			[][]byte{
				{0b_00100_000},
				{0b_01110_000},
				{0b_00010_000},
			},
			[][]bool{
				{false, false, true, false, false},
				{false, true, true, true, false},
				{false, false, false, true, false},
			},
		},
		{
			"Trivial Bigger",
			7, 5,
			[][]byte{
				{0b_0001000_0},
				{0b_0001000_0},
				{0b_0111110_0},
				{0b_0000010_0},
				{0b_0000010_0},
			},
			[][]bool{
				{false, false, false, true, false, false, false},
				{false, false, false, true, false, false, false},
				{false, true, true, true, true, true, false},
				{false, false, false, false, false, true, false},
				{false, false, false, false, false, true, false},
			},
		},
		{
			"Bigger Staggered",
			7, 5,
			[][]byte{
				{0b_0001000_0},
				{0b_0001000_0},
				{0b_0111110_0},
				{0b_0000100_0},
				{0b_0000100_0},
			},
			[][]bool{
				{false, false, false, true, false, false, false},
				{false, false, false, true, false, false, false},
				{false, true, true, true, true, true, false},
				{false, false, false, false, true, false, false},
				{false, false, false, false, true, false, false},
			},
		},
		{
			"Normal",
			11, 11,
			[][]byte{
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
			[][]bool{
				{false, false, false, false, false, true, false, false, false, false, false},
				{false, true, true, true, true, true, false, true, true, true, false},
				{false, false, false, false, false, true, false, false, false, true, false},
				{false, true, true, true, false, true, true, true, true, true, false},
				{false, true, false, true, false, false, false, false, false, true, false},
				{false, true, false, true, true, true, true, true, true, true, false},
				{false, false, false, true, false, false, false, true, false, true, false},
				{false, true, true, true, false, true, true, true, false, true, false},
				{false, true, false, false, false, false, false, false, false, true, false},
				{false, true, true, true, true, true, false, true, true, true, false},
				{false, false, false, false, false, true, false, false, false, false, false},
			},
		},
	}

	for _, test := range tests {
		rawMaze := RawMaze{
			Width:  test.width,
			Height: test.height,
			Data:   test.data,
		}

		for y, row := range test.expected {
			for x, expected := range row {
				if rawMaze.IsPath(x, y) != expected {
					t.Fatalf("%s: Wanted path at (%v, %v), apparently it isn't", test.name, x, y)
				}
			}
		}
	}
}
