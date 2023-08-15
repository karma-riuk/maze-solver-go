package solver

import (
	"maze-solver/maze"
	"slices"
)

type sorted_stack []*maze.Node

func (s *sorted_stack) insert(node *maze.Node, weights *map[*maze.Node]int) {
	var dummy *maze.Node
	*s = append(*s, dummy) // extend the slice

	i, _ := slices.BinarySearchFunc(*s, node, func(e, t *maze.Node) int {
		return (*weights)[t] - (*weights)[e]
	})

	copy((*s)[i+1:], (*s)[i:]) // make room
	(*s)[i] = node
}

func (s *sorted_stack) pop() *maze.Node {
	last_i := len(*s) - 1
	ret := (*s)[last_i]
	*s = (*s)[:last_i]
	return ret
}
