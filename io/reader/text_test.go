package reader

import (
	// "maze-solver/maze"
	// "maze-solver/utils"
	// "reflect"
	"testing"
)

func TestTextReadTrivial(t *testing.T) {
	/*
		tests := []struct {
			name     string
			filename string
			pathChar byte
			wallChar byte
			expected *maze.RawMaze
		}{
			{
				"Trivial",
				"../../assets/trivial.txt",
				' ',
				'#',
				&maze.RawMaze{
					PathChar: ' ',
					WallChar: '#',
					Data: []string{
						"## ##",
						"#   #",
						"### #",
					},
				},
			},
			{
				"Trivial Bigger",
				"../../assets/trivial-bigger.txt",
				' ',
				'#',
				&maze.RawMaze{
					PathChar: ' ',
					WallChar: '#',
					Data: []string{
						"### ###",
						"### ###",
						"#     #",
						"##### #",
						"##### #",
					},
				},
			},
			{
				"Bigger Staggered",
				"../../assets/trivial-bigger-staggered.txt",
				' ',
				'#',
				&maze.RawMaze{
					PathChar: ' ',
					WallChar: '#',
					Data: []string{
						"### ###",
						"### ###",
						"#     #",
						"#### ##",
						"#### ##",
					},
				},
			},
			{
				"Normal",
				"../../assets/normal.txt",
				' ',
				'#',
				&maze.RawMaze{
					PathChar: ' ',
					WallChar: '#',
					Data: []string{
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
	*/
}
