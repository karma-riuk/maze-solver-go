package maze

import "testing"

func TestRawMazeWall(t *testing.T) {
	tests := []struct {
		name     string
		pathChar byte
		wallChar byte
		data     []string
		expected [][]bool
	}{
		{
			"Trivial",
			' ',
			'#',
			[]string{
				"## ##",
				"#   #",
				"### #",
			},
			[][]bool{
				{true, true, false, true, true},
				{true, false, false, false, true},
				{true, true, true, false, true},
			},
		},
		{
			"Trivial Bigger",
			' ',
			'#',
			[]string{
				"### ###",
				"### ###",
				"#     #",
				"##### #",
				"##### #",
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
			' ',
			'#',
			[]string{
				"### ###",
				"### ###",
				"#     #",
				"#### ##",
				"#### ##",
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
			' ',
			'#',
			[]string{
				"##### #####",
				"#     #   #",
				"##### ### #",
				"#   #     #",
				"# # ##### #",
				"# #       #",
				"### ### # #",
				"#   #   # #",
				"# ####### #",
				"#     #   #",
				"##### #####",
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
			PathChar: test.pathChar,
			WallChar: test.wallChar,
			Data:     test.data,
		}

		for y, row := range test.expected {
			for x, expected := range row {
				if rawMaze.isWall(x, y) != expected {
					t.Fatalf("Wanted wall at (%v, %v), apparently it isn't", x, y)
				}
			}
		}
	}
}

func TestRawMazePath(t *testing.T) {
	tests := []struct {
		name     string
		pathChar byte
		wallChar byte
		data     []string
		expected [][]bool
	}{
		{
			"Trivial",
			' ',
			'#',
			[]string{
				"## ##",
				"#   #",
				"### #",
			},
			[][]bool{
				{false, false, true, false, false},
				{false, true, true, true, false},
				{false, false, false, true, false},
			},
		},
		{
			"Trivial Bigger",
			' ',
			'#',
			[]string{
				"### ###",
				"### ###",
				"#     #",
				"##### #",
				"##### #",
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
			' ',
			'#',
			[]string{
				"### ###",
				"### ###",
				"#     #",
				"#### ##",
				"#### ##",
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
			' ',
			'#',
			[]string{
				"##### #####",
				"#     #   #",
				"##### ### #",
				"#   #     #",
				"# # ##### #",
				"# #       #",
				"### ### # #",
				"#   #   # #",
				"# ####### #",
				"#     #   #",
				"##### #####",
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
			PathChar: test.pathChar,
			WallChar: test.wallChar,
			Data:     test.data,
		}

		for y, row := range test.expected {
			for x, expected := range row {
				if rawMaze.isPath(x, y) != expected {
					t.Fatalf("Wanted path at (%v, %v), apparently it isn't", x, y)
				}
			}
		}
	}
}
