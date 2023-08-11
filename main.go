package main

import (
	"errors"
	"fmt"
	"image/color"
	"maze-solver/io/reader"
	"maze-solver/io/writer"
	"maze-solver/maze/parser"
	"maze-solver/solver"
	"maze-solver/utils"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/mazznoer/colorgrad"
)

func main() {
	argparser := argparse.NewParser("maze-solver", "Solves the given maze (insane, right? who would've guessed?)")

	var verboseLevel *int = argparser.FlagCounter("v", "verbose", &argparse.Options{
		Help: `Verbose level of the solver
        0: nothing printed to stdout
        1: print the total time taken by the solver (time of the main() function)
        2: prints the time the solving algorithm took to run
        3: prints the time taken by each section (reader, solving algorithm, writer)`,
	})

	readerFactory := reader.ReaderFactory{}
	writerFactory := writer.WriterFactory{}
	solverFactory := solver.SolverFactory{}

	readerFactory.Type = reader.TYPES[".png"]
	readerFactory.Filename = argparser.String("i", "input", &argparse.Options{
		Help:    "Input file",
		Default: "maze.png",
		Validate: func(args []string) error {
			var ok bool
			extension := args[0][len(args[0])-4:]
			readerFactory.Type, ok = reader.TYPES[extension]
			if ok {
				return nil
			} else {
				return errors.New(fmt.Sprintf("Filetype not recognized %q", extension))
			}
		},
	})

	writerFactory.Type = writer.TYPES[".png"]
	writerFactory.Filename = argparser.String("o", "output", &argparse.Options{
		Help:    "Input file",
		Default: "maze_sol.png",
		Validate: func(args []string) error {
			var ok bool
			extension := args[0][len(args[0])-4:]
			writerFactory.Type, ok = writer.TYPES[extension]
			if ok {
				return nil
			} else {
				return errors.New(fmt.Sprintf("Filetype not recognized %q", extension))
			}
		},
	})

	readerFactory.PathChar = argparser.String("", "path-char-in", &argparse.Options{
		Help:    "Character to represent the path in a input text file",
		Default: " ",
		Validate: func(args []string) error {
			if len(args[0]) > 1 {
				return errors.New("Character must a string of length 1")
			}
			return nil
		},
	})

	readerFactory.WallChar = argparser.String("", "wall-char-in", &argparse.Options{
		Help:    "Character to represent the wall in a input text file",
		Default: "#",
		Validate: func(args []string) error {
			if len(args[0]) > 1 {
				return errors.New("Character must a string of length 1")
			}
			return nil
		},
	})

	writerFactory.PathChar = argparser.String("", "path-char-out", &argparse.Options{
		Help:    "Character to represent the path in a output text file",
		Default: " ",
		Validate: func(args []string) error {
			if len(args[0]) > 1 {
				return errors.New("Character must a string of length 1")
			}
			return nil
		},
	})

	writerFactory.WallChar = argparser.String("", "wall-char-out", &argparse.Options{
		Help:    "Character to represent the wall in a output text file",
		Default: "#",
		Validate: func(args []string) error {
			if len(args[0]) > 1 {
				return errors.New("Character must a string of length 1")
			}
			return nil
		},
	})

	cellSizeIn := argparser.Int("", "cell-size-in", &argparse.Options{
		Help:    "Size of a cell (in pixels) for input file of image type",
		Default: 20,
	})

	cellSizeOut := argparser.Int("", "cell-size-out", &argparse.Options{
		Help:    "Size of a cell (in pixels) for output file of image type",
		Default: 20,
	})

	solverFactory.Type = argparser.Selector("a", "algo", solver.TYPES, &argparse.Options{
		Help:    fmt.Sprintf("Algorithm to solve the maze, avaiable options: %s", strings.Join(solver.TYPES, ", ")),
		Default: solver.TYPES[0],
	})

	if err := argparser.Parse(os.Args); err != nil {
		fmt.Println(argparser.Usage(err))
		return
	}
	utils.VERBOSE_LEVEL = *verboseLevel

	readerFactory.CellHeight, readerFactory.CellWidth = cellSizeIn, cellSizeIn
	readerFactory.WallColor = color.RGBA{0, 0, 0, 255}
	readerFactory.PathColor = color.RGBA{255, 255, 255, 255}

	writerFactory.CellHeight, writerFactory.CellWidth = cellSizeOut, cellSizeOut
	writerFactory.WallColor = color.RGBA{0, 0, 0, 255}
	writerFactory.PathColor = color.RGBA{255, 255, 255, 255}
	writerFactory.SolutionGradient = colorgrad.Warm()

	defer utils.Timer("TOTAL", 1)()

	reader := readerFactory.Get()

	maze, err := parser.Parse(reader)
	utils.Check(err, "Couldn't read maze")

	solver := solverFactory.Get()
	solved := solver.Solve(maze)

	writer := writerFactory.Get(solved)

	err = writer.Write()
	utils.Check(err, "Couldn't write solved maze")
}
