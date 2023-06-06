package internal

type SearchInsertItem struct {
	ID       string
	Title    string
	Provider string
}

//counterfeiter:generate . Search

type Search interface {
	Query(term string) ([]string, error)
	Insert(*SearchInsertItem) error
	InsertBatch(items []*SearchInsertItem) error
	Clear() error
}
