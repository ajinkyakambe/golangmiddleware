package csv

type CSVContract interface {
	RowIterator(pos int) RowIterator
	Incorporate(other CSV)
	ToStringRFC4180() string
}

type RowIteratorContract interface {
	Get() []string
	Next() bool
	Pos() int
}
