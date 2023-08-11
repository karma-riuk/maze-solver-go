package reader

import (
	"bufio"
	"maze-solver/utils"
	"os"
)

type TextReader struct {
	Filename           string
	PathChar, WallChar byte
}

func (r *TextReader) Read() (*RawMaze, error) {
	defer utils.Timer("Text Reader", 3)()
	lines, err := getLines(r.Filename)
	if err != nil {
		return nil, err
	}

	strings_reader := StringsReader{
		PathChar: r.PathChar,
		WallChar: r.WallChar,
		Lines:    lines,
	}

	return strings_reader.Read()
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
