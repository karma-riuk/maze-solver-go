package maze

type Maze interface {
	maze()
}

type SolvedMaze interface {
	Maze
	solvedMaze()
}
