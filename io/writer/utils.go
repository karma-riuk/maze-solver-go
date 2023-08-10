package writer

import "maze-solver/maze"

func trivial() *maze.SolvedMaze {
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

	ret := &maze.SolvedMaze{
		Maze: &maze.Maze{
			Width:  5,
			Height: 3,
			Nodes:  nodes,
		},
		Solution: []*maze.Node{
			nodes[0], nodes[2], nodes[3], nodes[4],
		},
	}

	return ret
}

func bigger() *maze.SolvedMaze {
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

	ret := &maze.SolvedMaze{
		Maze: &maze.Maze{
			Width:  7,
			Height: 5,
			Nodes:  nodes,
		},
		Solution: []*maze.Node{
			nodes[0], nodes[2], nodes[3], nodes[4],
		},
	}

	return ret
}

func bigger_staggered() *maze.SolvedMaze {
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

	nodes[4].Left = nodes[2]
	nodes[4].Right = nodes[3]
	nodes[4].Down = nodes[5]

	nodes[5].Up = nodes[4]

	ret := &maze.SolvedMaze{
		Maze: &maze.Maze{
			Width:  7,
			Height: 5,
			Nodes:  nodes,
		},
		Solution: []*maze.Node{
			nodes[0], nodes[2], nodes[4], nodes[5],
		},
	}

	return ret
}

func normal() *maze.SolvedMaze {
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

	ret := &maze.SolvedMaze{
		Maze: &maze.Maze{
			Width:  11,
			Height: 11,
			Nodes:  nodes,
		},
		Solution: []*maze.Node{
			nodes[0],
			nodes[2],
			nodes[7],
			nodes[8],
			nodes[11],
			nodes[15],
			nodes[10],
			nodes[13],
			nodes[12],
			nodes[17],
			nodes[18],
			nodes[21],
		},
	}

	return ret
}
