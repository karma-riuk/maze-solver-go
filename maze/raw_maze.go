package maze

import (
	"strings"
)

const CHUNK_SIZE = 8 // size of a byte

type RawMaze struct {
	Width, Height int
	Data          [][]byte
}

func (m *RawMaze) String() string {
	var ret strings.Builder
	ret.WriteString("{\n")
	ret.WriteString("\tData: \n")
	for _, line := range m.Data {
		ret.WriteRune('\t')
		ret.WriteRune('\t')
		ret.Write(line) // TODO: prolly should fix this to make it readable
		ret.WriteRune('\n')
	}
	ret.WriteString("}")

	return ret.String()
}

func (m *RawMaze) IsPath(x int, y int) bool {
	chunk_index := x / CHUNK_SIZE
	chunk_rest := x % CHUNK_SIZE
	chunk := m.Data[y][chunk_index]
	return chunk&(1<<(CHUNK_SIZE-1-chunk_rest)) != 0
}

func (m *RawMaze) IsWall(x int, y int) bool {
	chunk_index := x / CHUNK_SIZE
	chunk_rest := x % CHUNK_SIZE
	chunk := m.Data[y][chunk_index]
	return chunk&(1<<(CHUNK_SIZE-1-chunk_rest)) == 0
}
