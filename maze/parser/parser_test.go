package parser

import (
	"fmt"
	"maze-solver/io/reader"
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

	reader := reader.TextReader{
		Filename: "../../assets/trivial.txt",
		PathChar: ' ',
		WallChar: '#',
	}

	got, err := Parse(reader)
	utils.Check(err, "Couldn't create maze from %q", reader.Filename)

	utils.AssertEqual(t, got.Width, 5, "Normal: width differ")
	utils.AssertEqual(t, got.Height, 3, "Normal: height differ")

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

	reader := reader.TextReader{
		Filename: "../../assets/trivial-bigger.txt",
		PathChar: ' ',
		WallChar: '#',
	}

	got, err := Parse(reader)
	utils.Check(err, "Couldn't create maze from %q", reader.Filename)

	utils.AssertEqual(t, got.Width, 7, "Normal: width differ")
	utils.AssertEqual(t, got.Height, 5, "Normal: height differ")

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

	reader := reader.TextReader{
		Filename: "../../assets/trivial-bigger-staggered.txt",
		PathChar: ' ',
		WallChar: '#',
	}

	got, err := Parse(reader)
	utils.Check(err, "Couldn't create maze from %q", reader.Filename)

	utils.AssertEqual(t, got.Width, 7, "Normal: width differ")
	utils.AssertEqual(t, got.Height, 5, "Normal: height differ")

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

func TestTextReadNormal(t *testing.T) {
	/* normal.txt
		##### #####
		#     #   #
		##### ### #
		#   #     #
		# # ##### #
		# #       #
		### ### # #
		#   #   # #
		# ####### #
		#     #   #
		##### #####

	    Nodes are
		#####0#####
		#1   2#3 4#
		##### ### #
		#5 6#7   8#
		# # ##### #
		#9#A   F B#
		### ### # #
		#C D#E G# #
		# ####### #
		#H   I#J K#
		#####L#####
	*/
	nodes := make([]*maze.Node, 22)

	// ---- Node creation ----
	coords := []struct{ x, y int }{
		{5, 0}, // 0

		{1, 1}, // 1
		{5, 1}, // 2
		{7, 1}, // 3
		{9, 1}, // 4

		{1, 3}, // 5
		{3, 3}, // 6
		{5, 3}, // 7
		{9, 3}, // 8

		{1, 5}, // 9
		{3, 5}, // A (10)
		{9, 5}, // B (11)

		{1, 7}, // C (12)
		{3, 7}, // D (13)
		{5, 7}, // E (14)
		{7, 5}, // F (15)
		{7, 7}, // G (16)

		{1, 9}, // H (17)
		{5, 9}, // I (18)
		{7, 9}, // J (19)
		{9, 9}, // K (20)

		{5, 10}, // L (21)
	}

	for i, coord := range coords {
		nodes[i] = maze.NewNode(maze.Coordinates{X: coord.x, Y: coord.y})
	}

	// ---- Node linking ----
	// Vertical
	links := []struct {
		from, to int
	}{
		{0, 2},
		{2, 7},
		{4, 8},
		{5, 9},
		{6, 10},
		{8, 11},
		{10, 13},
		{15, 16},
		{11, 20},
		{12, 17},
		{18, 21},
	}
	for _, link := range links {
		nodes[link.from].Down = nodes[link.to]
		nodes[link.to].Up = nodes[link.from]
	}

	links = []struct {
		from, to int
	}{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
		{10, 15},
		{15, 11},
		{12, 13},
		{14, 16},
		{17, 18},
		{19, 20},
	}
	for _, link := range links {
		nodes[link.from].Right = nodes[link.to]
		nodes[link.to].Left = nodes[link.from]
	}

	reader := reader.TextReader{
		Filename: "../../assets/normal.txt",
		PathChar: ' ',
		WallChar: '#',
	}

	got, err := Parse(reader)
	utils.Check(err, "Couldn't create maze from %q", reader.Filename)

	utils.AssertEqual(t, got.Width, 11, "Normal: width differ")
	utils.AssertEqual(t, got.Height, 11, "Normal: height differ")

	if len(nodes) != len(got.Nodes) {
		for i, node := range got.Nodes {
			fmt.Printf("%v: %v\n", i, node)
		}
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
