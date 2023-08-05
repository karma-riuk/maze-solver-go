package reader

type Reader interface {
	Read() (*RawMaze, error)
}
