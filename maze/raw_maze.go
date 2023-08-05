package maze

import (
	"fmt"
	"strings"
)

type RawMaze struct {
	PathChar, WallChar byte
	Data               []string
}

func (m *RawMaze) String() string {
	var ret strings.Builder
	ret.WriteString("{\n")
	ret.WriteString(fmt.Sprintf("\tPathChar: %v,\n", m.PathChar))
	ret.WriteString(fmt.Sprintf("\tWallChar: %v,\n", m.WallChar))
	ret.WriteString("\tData: \n")
	for _, line := range m.Data {
		ret.WriteRune('\t')
		ret.WriteRune('\t')
		ret.WriteString(line)
		ret.WriteRune('\n')
	}
	ret.WriteString("}")

	return ret.String()
}

func (m *RawMaze) isPath(x int, y int) bool {
	return m.Data[y][x] == m.PathChar
}

func (m *RawMaze) isWall(x int, y int) bool {
	return m.Data[y][x] == m.WallChar
}
