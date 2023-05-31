package internal

import (
	"github.com/blevesearch/bleve"
)

type searchBleve struct {
	Index bleve.Index
}

func NewSearchBleve(index bleve.Index) Search {
	return &searchBleve{index}
}

func (search *searchBleve) Query(term string) ([]string, error) {
	q := bleve.NewFuzzyQuery(term)
	req := bleve.NewSearchRequest(q)
	res, err := search.Index.Search(req)
	if err != nil {
		return nil, err
	}
	var result []string
	for i := range res.Hits {
		id := res.Hits[i].ID
		result = append(result, id)
	}
	return result, nil
}

func (search *searchBleve) Insert(id string, text string) error {
	return search.InsertBatch([]*SearchInsertItem{{id, text}})
}

func (search *searchBleve) InsertBatch(items []*SearchInsertItem) error {
	b := search.Index.NewBatch()
	for i := range items {
		if err := b.Index(items[i].ID, items); err != nil {
			return err
		}
	}
	return search.Index.Batch(b)
}
