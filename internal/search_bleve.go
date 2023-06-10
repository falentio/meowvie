package internal

import (
	"github.com/blevesearch/bleve"
)

var _ Search = new(searchBleve)

type searchBleve struct {
	Index bleve.Index
}

func NewSearchBleve(index bleve.Index) Search {
	return &searchBleve{index}
}

func (search *searchBleve) Query(term string) ([]string, error) {
	q := bleve.NewMatchQuery(term)
	q.SetFuzziness(1)
	q.SetBoost(1.9)
	qq := bleve.NewQueryStringQuery(term)
	qq.SetBoost(0.8)
	qqq := bleve.NewDisjunctionQuery(q, qq)
	req := bleve.NewSearchRequest(qqq)
	req.SortBy([]string{"-_score", "-_id"})
	res, err := search.Index.Search(req)
	if err != nil {
		return nil, err
	}
	var result []string
	for i := range res.Hits {
		if i > 9 {
			break
		}
		if res.Hits[i].Score < 0.2 {
			break
		}
		id := res.Hits[i].ID
		result = append(result, id)
	}
	return result, nil
}

func (search *searchBleve) Insert(item *SearchInsertItem) error {
	return search.InsertBatch([]*SearchInsertItem{item})
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

func (search *searchBleve) Clear() error {
	req := bleve.NewSearchRequest(bleve.NewMatchAllQuery())
	res, err := search.Index.Search(req)
	if err != nil {
		return err
	}
	b := search.Index.NewBatch()
	for _, h := range res.Hits {
		b.Delete(h.ID)
	}
	return search.Index.Batch(b)
}
