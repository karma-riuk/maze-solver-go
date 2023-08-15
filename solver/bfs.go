package solver

import (
	"errors"
	"fmt"
	"maze-solver/maze"
	"maze-solver/utils"
	"strings"
)

type BFSSolver struct {
	queue *Queue
}

type Queue struct {
	head, tail *Element
}

type Element struct {
	prev, next *Element
	value      []*maze.Node
}

func (q *Queue) enqueue(v []*maze.Node) {
	prev_last := q.tail
	new_elem := &Element{
		prev:  prev_last,
		next:  nil,
		value: v,
	}
	if prev_last != nil {
		prev_last.next = new_elem
	}

	q.tail = new_elem

	if q.head == nil {
		q.head = new_elem
	}
}

func (q *Queue) dequeue() ([]*maze.Node, error) {
	if q.head == nil {
		return nil, errors.New("Can't dequeue and empty queue")
	}
	ret := q.head.value
	q.head = q.head.next
	if q.head != nil {
		q.head.prev = nil
	} else {
		q.tail = nil
	}

	return ret, nil
}

func (q Queue) String() string {
	var ret strings.Builder
	i := 0
	for history := q.head; history != nil; history = history.next {
		ret.WriteString(fmt.Sprintf("%v: %v\n", i, history_str(history.value)))
		i++
	}
	return ret.String()
}

func history_str(history []*maze.Node) string {
	var ret strings.Builder
	for _, node := range history {
		ret.WriteString(fmt.Sprintf("%v ", node.Coords))
	}
	return ret.String()
}

func (s *BFSSolver) Solve(m *maze.Maze) *maze.SolvedMaze {
	defer utils.Timer("BFS algorithm", 2)()

	current, end := m.Nodes[0], m.Nodes[len(m.Nodes)-1]
	s.queue = &Queue{
		head: nil,
		tail: nil,
	}

	current_history := make([]*maze.Node, 0, len(m.Nodes))
	current_history = append(current_history, current)

	var err error
	for current != end {
		current.Visited = true

		s.addIfNotVisited(current.Down, current_history)
		s.addIfNotVisited(current.Left, current_history)
		s.addIfNotVisited(current.Right, current_history)
		s.addIfNotVisited(current.Up, current_history)

		current_history, err = s.queue.dequeue()
		if err != nil {
			panic(err)
		}
		current = current_history[len(current_history)-1]
	}

	return &maze.SolvedMaze{
		Maze:     m,
		Solution: current_history,
	}
}

func (s *BFSSolver) addIfNotVisited(node *maze.Node, current_history []*maze.Node) {
	if !wasVisited(node) {
		new_history := make([]*maze.Node, len(current_history)+1)
		copy(new_history, current_history)
		new_history[len(current_history)] = node

		s.queue.enqueue(new_history)
	}
}
