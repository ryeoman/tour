package database

type Read struct {
}

type ReadAdapter interface {
	GetById()
}
