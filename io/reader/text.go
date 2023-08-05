package reader

import (
	"bufio"
	"maze-solver/maze"
	"os"
)

type TextReader struct {
	Filename           string
	PathChar, WallChar byte
}

func (r TextReader) Read() (*maze.RawMaze, error) {
	lines, err := getLines(r.Filename)
	if err != nil {
		return nil, err
	}

	return &maze.RawMaze{
		PathChar: r.PathChar,
		WallChar: r.WallChar,
		Data:     *lines,
	}, nil
}

func getLines(filename string) (*[]string, error) {
	var lines []string
	if _, err := os.Stat(filename); err != nil {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return &lines, nil
}
