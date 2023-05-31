package internal

type SearchInsertItem struct {
	ID   string
	Text string
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Search

type Search interface {
	Query(term string) ([]string, error)
	Insert(id string, text string) error
	InsertBatch(items []*SearchInsertItem) error
}
