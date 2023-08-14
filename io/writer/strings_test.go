package writer

import (
	"maze-solver/maze"
	"maze-solver/utils"
	"testing"
)

func TestStringsWriter(t *testing.T) {
	tests := []struct {
		name                             string
		m                                *maze.SolvedMaze
		pathChar, wallChar, solutionChar byte
		expected                         []string
	}{
		{
			"Trivial",
			trivial(),
			' ', '#', '.',
			[]string{
				"##.##",
				"# ..#",
				"###.#",
			},
		},
		{
			"Bigger",
			bigger(),
			'_', '~', '*',
			[]string{
				"~~~*~~~",
				"~~~*~~~",
				"~__***~",
				"~~~~~*~",
				"~~~~~*~",
			},
		},
		{
			"Bigger Staggered",
			bigger_staggered(),
			' ', '#', '.',
			[]string{
				"###.###",
				"###.###",
				"#  .. #",
				"####.##",
				"####.##",
			},
		},
		{
			"Normal",
			normal(),
			' ', '#', '.',
			[]string{
				"#####.#####",
				"#    .#   #",
				"#####.### #",
				"#   #.....#",
				"# # #####.#",
				"# #.......#",
				"###.### # #",
				"#...#   # #",
				"#.####### #",
				"#.....#   #",
				"#####.#####",
			},
		},
	}

	for _, test := range tests {
		writer := StringsWriter{
			PathChar:     test.pathChar,
			WallChar:     test.wallChar,
			SolutionChar: test.solutionChar,
			Maze:         test.m,
		}
		writer.Write()
		got := writer.GetLines()

		utils.AssertEqual(t, len(got), len(test.expected), "%s: different amount of lines.", test.name)

		for i, line := range test.expected {
			utils.AssertEqual(t, got[i], line, "%s, line %v: not what we expected.", test.name, i)
		}
	}
}
