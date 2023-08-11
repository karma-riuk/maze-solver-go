package writer

import (
	"testing"
)

func TestImageWriter(t *testing.T) {
	// pathGradient, err := colorgrad.NewGradient().Colors(color.White).Build()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// tests := []struct {
	// 	name                  string
	// 	filename              string
	// 	m                     *maze.SolvedMaze
	// 	CellWidth, cellHeight int
	// 	pathColor, wallColor  color.Color
	// 	gradient              colorgrad.Gradient
	// }{
	// 	{
	// 		"Trivial",
	// 		"../../out/trivial_sol.png",
	// 		trivial(),
	// 		40, 40,
	// 		color.White, color.Black,
	// 		colorgrad.Warm(),
	// 	},
	// 	{
	// 		"Bigger",
	// 		"../../out/bigger_sol.png",
	// 		bigger(),
	// 		40, 40,
	// 		color.White, color.Black,
	// 		colorgrad.Warm(),
	// 	},
	// 	{
	// 		"Bigger Staggered",
	// 		"../../out/bigger_staggered_sol.png",
	// 		bigger_staggered(),
	// 		40, 40,
	// 		color.White, color.Black,
	// 		pathGradient,
	// 	},
	// 	{
	// 		"Normal",
	// 		"../../out/normal_sol.png",
	// 		normal(),
	// 		40, 40,
	// 		color.White, color.Black,
	// 		colorgrad.Warm(),
	// 	},
	// }
	//
	// for _, test := range tests {
	// 	writer := ImageWriter{
	// 		Filename:         test.filename,
	// 		Maze:             test.m,
	// 		CellWidth:        test.CellWidth,
	// 		CellHeight:       test.cellHeight,
	// 		WallColor:        test.wallColor,
	// 		PathColor:        test.pathColor,
	// 		SolutionGradient: test.gradient,
	// 	}
	//
	// 	err := writer.Write()
	// 	if err != nil {
	// 		t.Fatalf("%s: couldn't write solution, got following error\n%v", test.name, err)
	// 	}
	// }
}
