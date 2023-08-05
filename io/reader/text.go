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
	var lines []string

	if _, err := os.Stat(r.Filename); err != nil {
		return nil, err
	}

	file, err := os.Open(r.Filename)
	if err != nil {
		return nil, err
	}

	{
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			line := scanner.Text()
			lines = append(lines, line)
		}
		file.Close()
	}

	return &maze.RawMaze{PathChar: r.PathChar, WallChar: r.WallChar, Data: lines}, nil
}
