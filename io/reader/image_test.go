package reader

import (
	"image/color"
	"maze-solver/utils"
	"testing"
)

func TestImageReader(t *testing.T) {
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	tests := []struct {
		name                  string
		width, height         int
		cellWidth, cellHeight int
		pathColor, wallColor  color.Color
		filename              string
		expected              [][]byte
	}{
		{
			"Trivial",
			5, 3,
			40, 40,
			white, black,
			"../../assets/trivial.png",
			[][]byte{
				{0b_00100_000},
				{0b_01110_000},
				{0b_00010_000},
			},
		},
		{
			"Trivial Bigger",
			7, 5,
			40, 40,
			white, black,
			"../../assets/trivial-bigger.png",
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
			40, 40,

			white, black,
			"../../assets/trivial-bigger-staggered.png",
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
			40, 40,
			white, black,
			"../../assets/normal.png",
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
		{
			"Normal2",
			15, 15,
			20, 20,
			white, black,
			"../../assets/normal2.png",
			[][]byte{
				{0b00000001, 0b0000000_0},
				{0b01110111, 0b1111010_0},
				{0b00010100, 0b0001010_0},
				{0b01110111, 0b1101110_0},
				{0b01000000, 0b0000010_0},
				{0b01111101, 0b1111010_0},
				{0b00010101, 0b0000010_0},
				{0b01110111, 0b0111010_0},
				{0b01000100, 0b0101010_0},
				{0b01011101, 0b1101110_0},
				{0b01000101, 0b0000000_0},
				{0b01110101, 0b0111110_0},
				{0b01010001, 0b0001010_0},
				{0b01011111, 0b1111010_0},
				{0b00000001, 0b0000000_0},
			},
		},
	}

	for _, test := range tests {
		reader := ImageReader{
			Filename:   test.filename,
			PathColor:  test.pathColor,
			WallColor:  test.wallColor,
			CellWidth:  test.cellWidth,
			CellHeight: test.cellHeight,
		}
		got, err := reader.Read()
		if err != nil {
			t.Fatalf("%s: got error while reading, got\n%v", test.filename, err)
		}

		utils.AssertEqual(t, got.Width, test.width, "%s: width of raw maze don't match", test.name)
		utils.AssertEqual(t, got.Height, test.height, "%s: height of raw maze don't match", test.name)
		utils.AssertEqual(t, len(got.Data), len(test.expected), "%s: don't have the same number of rows", test.name)

		for y, line_exp := range test.expected {
			line_got := got.Data[y]
			utils.AssertEqual(t, len(line_got), len(line_exp), "%s (line %v): don't have same number of chunks", test.name, y)

			for i, chunk_exp := range line_exp {
				chunk_got := line_got[i]
				if chunk_got != chunk_exp {
					t.Fatalf("%s (line %v): chunk %v don't coincide, %08b, want %08b", test.name, y, i, chunk_got, chunk_exp)
				}
			}
		}
	}
}
