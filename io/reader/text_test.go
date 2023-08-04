package reader

import (
	"fmt"
	"maze-solver/maze"
	"maze-solver/utils"
	"testing"
)

func TestTextReadTrivial(t *testing.T) {
	/* trivial.txt
	   ## ##
	   #   #
	   ### #

	   Nodes are
	   ##0##
	   #123#
	   ###4#
	*/
	nodes := make([]*maze.Node, 5)

	nodes[0] = maze.NewNode(maze.Coordinates{X: 2, Y: 0})

	nodes[1] = maze.NewNode(maze.Coordinates{X: 1, Y: 1})
	nodes[2] = maze.NewNode(maze.Coordinates{X: 2, Y: 1})
	nodes[3] = maze.NewNode(maze.Coordinates{X: 3, Y: 1})

	nodes[4] = maze.NewNode(maze.Coordinates{X: 3, Y: 2})

	nodes[0].Down = nodes[2]

	nodes[1].Right = nodes[2]

	nodes[2].Up = nodes[0]
	nodes[2].Left = nodes[1]
	nodes[2].Right = nodes[3]

	nodes[3].Left = nodes[2]
	nodes[3].Down = nodes[4]

	nodes[4].Up = nodes[3]

	reader := TextReader{
		PathChar: ' ',
		WallChar: '#',
	}

	filename := "../../assets/trivial.txt"
	got, err := reader.Read(filename)
	utils.Check(err, "Couldn't create maze from %q", filename)

	if len(nodes) != len(got.Nodes) {
		t.Fatalf("Didn't get the same size of nodes: %v, want %v", len(got.Nodes), len(nodes))
	}

	for i, got := range got.Nodes {
		expected := nodes[i]

		checkNode(t, i, got, expected, "")
		checkNode(t, i, got.Left, expected.Left, "left")
		checkNode(t, i, got.Right, expected.Right, "Right")
		checkNode(t, i, got.Up, expected.Up, "Up")
		checkNode(t, i, got.Down, expected.Down, "Down")
	}
}

func TestTextReadTrivialBigger(t *testing.T) {
	/* trivial-bigger.txt
	   ### ###
	   ### ###
	   #     #
	   ##### #
	   ##### #

	   Nodes are
	   ###0###
	   ### ###
	   #1 2 3#
	   ##### #
	   #####4#
	*/
	nodes := make([]*maze.Node, 5)

	nodes[0] = maze.NewNode(maze.Coordinates{X: 3, Y: 0})

	nodes[1] = maze.NewNode(maze.Coordinates{X: 1, Y: 2})
	nodes[2] = maze.NewNode(maze.Coordinates{X: 3, Y: 2})
	nodes[3] = maze.NewNode(maze.Coordinates{X: 5, Y: 2})

	nodes[4] = maze.NewNode(maze.Coordinates{X: 5, Y: 4})

	nodes[0].Down = nodes[2]

	nodes[1].Right = nodes[2]

	nodes[2].Up = nodes[0]
	nodes[2].Left = nodes[1]
	nodes[2].Right = nodes[3]

	nodes[3].Left = nodes[2]
	nodes[3].Down = nodes[4]

	nodes[4].Up = nodes[3]

	reader := TextReader{
		PathChar: ' ',
		WallChar: '#',
	}

	filename := "../../assets/trivial-bigger.txt"
	got, err := reader.Read(filename)
	utils.Check(err, "Couldn't create maze from %q", filename)

	if len(nodes) != len(got.Nodes) {
		t.Fatalf("Didn't get the same size of nodes: %v, want %v", len(got.Nodes), len(nodes))
	}

	for _, node := range got.Nodes {
		fmt.Println(node)
	}

	for i, got := range got.Nodes {
		expected := nodes[i]

		checkNode(t, i, got, expected, "")
		checkNode(t, i, got.Left, expected.Left, "left")
		checkNode(t, i, got.Right, expected.Right, "Right")
		checkNode(t, i, got.Up, expected.Up, "Up")
		checkNode(t, i, got.Down, expected.Down, "Down")
	}
}

func TestTextReadTrivialBiggerStaggered(t *testing.T) {
	/* trivial-bigger-staggered.txt
	   ### ###
	   ### ###
	   #     #
	   #### ##
	   #### ##

	   Nodes are
	   ###0###
	   ### ###
	   #1 243#
	   #### ##
	   ####5##
	*/
	fmt.Println("---------- STAGGERED ----------")
	nodes := make([]*maze.Node, 6)

	nodes[0] = maze.NewNode(maze.Coordinates{X: 3, Y: 0})

	nodes[1] = maze.NewNode(maze.Coordinates{X: 1, Y: 2})
	nodes[2] = maze.NewNode(maze.Coordinates{X: 3, Y: 2})
	nodes[3] = maze.NewNode(maze.Coordinates{X: 5, Y: 2})

	nodes[4] = maze.NewNode(maze.Coordinates{X: 4, Y: 2})

	nodes[5] = maze.NewNode(maze.Coordinates{X: 4, Y: 4})

	nodes[0].Down = nodes[2]

	nodes[1].Right = nodes[2]

	nodes[2].Up = nodes[0]
	nodes[2].Left = nodes[1]
	nodes[2].Right = nodes[4]

	nodes[3].Left = nodes[4]

	nodes[4].Down = nodes[5]
	nodes[4].Left = nodes[2]
	nodes[4].Right = nodes[3]

	nodes[5].Up = nodes[4]

	reader := TextReader{
		PathChar: ' ',
		WallChar: '#',
	}

	filename := "../../assets/trivial-bigger-staggered.txt"
	got, err := reader.Read(filename)
	utils.Check(err, "Couldn't create maze from %q", filename)

	if len(nodes) != len(got.Nodes) {
		t.Fatalf("Didn't get the same size of nodes: %v, want %v", len(got.Nodes), len(nodes))
	}

	for _, node := range got.Nodes {
		fmt.Println(node)
	}

	for i, got := range got.Nodes {
		expected := nodes[i]

		checkNode(t, i, got, expected, "")
		checkNode(t, i, got.Left, expected.Left, "left")
		checkNode(t, i, got.Right, expected.Right, "Right")
		checkNode(t, i, got.Up, expected.Up, "Up")
		checkNode(t, i, got.Down, expected.Down, "Down")
	}
}

func checkNode(t *testing.T, i int, got *maze.Node, expected *maze.Node, side string) {
	if expected == nil {
		return
	}

	if got == nil {
		t.Fatalf("No %s node of %v, want %v", side, i, expected.Coords)
	}

	if got.Coords != expected.Coords {
		t.Fatalf("Coords %s node of %v: %v, but want %v", side, i, got.Coords, expected.Coords)
	}
}
