package reader

import (
	"fmt"
	"maze-solver/utils"
)

type StringsReader struct {
	PathChar, WallChar byte
	Lines              *[]string
}

func (r *StringsReader) Read() (*RawMaze, error) {
	defer utils.Timer("Strings Reader", 3)()
	width, height := len((*r.Lines)[0]), len(*r.Lines)
	ret := &RawMaze{
		Width:  width,
		Height: height,
		Data:   make([][]byte, height),
	}

	for i := 0; i < height; i++ {
		ret.Data[i] = make([]byte, width/CHUNK_SIZE+1)
	}

	for y, line := range *r.Lines {
		r.processLine(line, &ret.Data[y])
	}

	return ret, nil
}

func (r *StringsReader) processLine(line string, dest *[]byte) {
	n_chunks := len(line)/CHUNK_SIZE + 1

	if len(*dest) != n_chunks {
		panic(fmt.Sprintf("The row that should receive the chunks does not have the correct length (%v, want %v)", len(*dest), n_chunks))
	}

	for i := 0; i < n_chunks; i++ {
		var chunk byte = 0 // all walls

		end_index := utils.Min((i+1)*CHUNK_SIZE, len(line))

		for x, c := range line[i*CHUNK_SIZE : end_index] {
			if c == rune(r.PathChar) {
				chunk |= 1 << (CHUNK_SIZE - 1 - x)
			}
		}

		(*dest)[i] = chunk
	}
}
