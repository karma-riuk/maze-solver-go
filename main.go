package main

import (
	"maze-solver/io/reader"
	"maze-solver/io/writer"
	"maze-solver/solver"
	"maze-solver/utils"
)

func main() {
	input := "filename"
	output := "filename"

	reader := &reader.TextReader{PathChar: ' ', WallChar: '#'}
	writer := &writer.ImageWriter{}

	solver := &solver.Bfs{}

	maze, err := reader.Read(input)
	utils.Check(err, "Couldn't read maze from %q", input)

	solved := solver.Solve(maze)
	err = writer.Write(output, solved)
	utils.Check(err, "Couldn't write solved maze to %q", output)
}
