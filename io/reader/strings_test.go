package reader

import (
	"testing"
)

func TestStringsReader(t *testing.T) {
	tests := []struct {
		name          string
		width, height int
		pathChar      byte
		wallChar      byte
		lines         []string
		expected      [][]byte
	}{
		{
			"Trivial",
			5, 3,
			' ',
			'#',
			[]string{
				"## ##",
				"#   #",
				"### #",
			},
			[][]byte{
				{0b_00100_000},
				{0b_01110_000},
				{0b_00010_000},
			},
		},
		{
			"Trivial Bigger",
			7, 5,
			' ',
			'#',
			[]string{
				"### ###",
				"### ###",
				"#     #",
				"##### #",
				"##### #",
			},
			[][]byte{
				{0b_0001000_0},
				{0b_0001000_0},
				{0b_0111110_0},
				{0b_0000010_0},
				{0b_0000010_0},
			},
		},
		{
			"Bigger Staggered",
			7, 5,
			' ',
			'#',
			[]string{
				"### ###",
				"### ###",
				"#     #",
				"#### ##",
				"#### ##",
			},
			[][]byte{
				{0b_0001000_0},
				{0b_0001000_0},
				{0b_0111110_0},
				{0b_0000100_0},
				{0b_0000100_0},
			},
		},
		{
			"Normal",
			11, 11,
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
		},
	}

	for _, test := range tests {
		reader := StringsReader{
			PathChar: test.pathChar,
			WallChar: test.wallChar,
			Lines:    &test.lines,
		}
		got, _ := reader.Read()

		assertEqual(t, got.Width, test.width, "%s: width of raw maze don't match", test.name)
		assertEqual(t, got.Height, test.height, "%s: height of raw maze don't match", test.name)
		assertEqual(t, len(got.Data), len(test.expected), "%s: don't have the same number of rows", test.name)

		for y, line_exp := range test.expected {
			line_got := got.Data[y]
			assertEqual(t, len(line_got), len(line_exp), "%s (line %v): don't have same number of chunks, %v, want %v", test.name, y)

			for i, chunk_exp := range line_exp {
				chunk_got := line_got[i]
				if chunk_got != chunk_exp {
					t.Fatalf("%s (line %v): chunk %v don't coincide, %08b, want %08b", test.name, y, i, chunk_got, chunk_exp)
				}
			}
		}
	}
}

func assertEqual[T comparable](t *testing.T, got T, want T, msg string, args ...any) {
	args = append(args, got, want)
	if got != want {
		t.Fatalf(msg+"\nGot: %v, Want: %v", args...)
	}
}
