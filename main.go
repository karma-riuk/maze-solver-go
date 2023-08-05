package main

import (
	"maze-solver/io/reader"
	"maze-solver/io/writer"
	"maze-solver/maze/parser"
	"maze-solver/solver"
	"maze-solver/utils"
)

func main() {
	output := "filename"

	reader := &reader.TextReader{Filename: "filename", PathChar: ' ', WallChar: '#'}
	writer := &writer.ImageWriter{}

	solver := &solver.Bfs{}

	maze, err := parser.Parse(reader)
	utils.Check(err, "Couldn't read maze from %q", reader.Filename)

	solved := solver.Solve(maze)
	err = writer.Write(output, solved)
	utils.Check(err, "Couldn't write solved maze to %q", output)
}
