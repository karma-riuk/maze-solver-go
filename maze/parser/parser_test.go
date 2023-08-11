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

	reader := &reader.TextReader{
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

	reader := &reader.TextReader{
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

	reader := &reader.TextReader{
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

	reader := &reader.TextReader{
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

func TestTextReadNormal2(t *testing.T) {
	/* normal2.txt
	   ####### #######
	   #   #       # #
	   ### # ##### # #
	   #   #     #   #
	   # ########### #
	   #     #     # #
	   ### # # ##### #
	   #   #   #   # #
	   # ### ### # # #
	   # #   #   #   #
	   # ### # #######
	   #   # # #     #
	   # # ### ### # #
	   # #         # #
	   ####### #######

	   Nodes are
	   #######0#######
	   #1 2#3 4   5#B#
	   ### # ##### # #
	   #6 7#8   9#A C#
	   # ########### #
	   #D I E#F   G# #
	   ### # # ##### #
	   #H J#K L#M N# #
	   # ### ### # # #
	   # #O P#Q R#S T#
	   # ### # #######
	   #U V#W# #X C Y#
	   # # ### ### # #
	   #Z#A   B   D#E#
	   #######F#######
	*/
	nodes := make([]*maze.Node, 42)

	// ---- Node creation ----
	coords := []struct{ x, y int }{
		{7, 0}, // 0

		{1, 1},  // 1
		{3, 1},  // 2
		{5, 1},  // 3
		{7, 1},  // 4
		{11, 1}, // 5

		{1, 3},  // 6
		{3, 3},  // 7
		{5, 3},  // 8
		{9, 3},  // 9
		{11, 3}, // 10 (A)
		{13, 1}, // 11 (B)
		{13, 3}, // 12 (C)

		{1, 5},  // 13 (D)
		{5, 5},  // 14 (E)
		{7, 5},  // 15 (F)
		{11, 5}, // 16 (G)

		{1, 7},  // 17 (H)
		{3, 5},  // 18 (I)
		{3, 7},  // 19 (J)
		{5, 7},  // 20 (K)
		{7, 7},  // 21 (L)
		{9, 7},  // 22 (M)
		{11, 7}, // 23 (N)

		{3, 9},  // 24 (O)
		{5, 9},  // 25 (P)
		{7, 9},  // 26 (Q)
		{9, 9},  // 27 (R)
		{11, 9}, // 28 (S)
		{13, 9}, // 29 (T)

		{1, 11},  // 30 (U)
		{3, 11},  // 31 (V)
		{5, 11},  // 32 (W)
		{9, 11},  // 33 (X)
		{13, 11}, // 34 (Y)

		{1, 13},  // 35 (Z)
		{3, 13},  // 36 (AA)
		{7, 13},  // 37 (AB)
		{11, 11}, // 38 (AC)
		{11, 13}, // 39 (AD)
		{13, 13}, // 40 (AE)

		{7, 14}, // 42 (AF)
	}

	for i, coord := range coords {
		nodes[i] = maze.NewNode(maze.Coordinates{X: coord.x, Y: coord.y})
	}

	// ---- Node linking ----
	// Vertical
	links := []struct {
		from, to int
	}{
		{0, 4},
		{2, 7},
		{3, 8},
		{5, 10},
		{11, 12},
		{6, 13},
		{12, 29},
		{18, 19},
		{14, 20},
		{15, 21},
		{17, 30},
		{20, 25},
		{22, 27},
		{23, 28},
		{25, 32},
		{26, 37},
		{30, 35},
		{31, 36},
		{38, 39},
		{34, 40},
		{37, 41},
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
		{4, 5},

		{6, 7},
		{8, 9},
		{10, 12},

		{13, 18},
		{18, 14},
		{15, 16},

		{17, 19},
		{20, 21},
		{22, 23},

		{24, 25},
		{26, 27},
		{28, 29},

		{30, 31},
		{33, 38},
		{38, 34},
		{36, 37},
		{37, 39},
	}
	for _, link := range links {
		nodes[link.from].Right = nodes[link.to]
		nodes[link.to].Left = nodes[link.from]
	}

	reader := &reader.TextReader{
		Filename: "../../assets/normal2.txt",
		PathChar: ' ',
		WallChar: '#',
	}

	got, err := Parse(reader)
	utils.Check(err, "Couldn't create maze from %q", reader.Filename)

	utils.AssertEqual(t, got.Width, 15, "Normal 2: width differ")
	utils.AssertEqual(t, got.Height, 15, "Normal 2: height differ")

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
	if expected == nil && got != nil {
		t.Fatalf("Somehow there is a node %s of %v, didn't want any", side, i)
	}

	if expected == nil && got == nil {
		return
	}

	if expected != nil && got == nil {
		t.Fatalf("No %s node of %v, want %v", side, i, expected.Coords)
	}

	if got.Coords != expected.Coords {
		t.Fatalf("Coords %s node of %v: %v, but want %v", side, i, got.Coords, expected.Coords)
	}
}
