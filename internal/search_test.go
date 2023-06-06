package internal_test

import (
	"testing"

	. "meowvie/internal"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	index, err := bleve.NewMemOnly(mapping.NewIndexMapping())
	if err != nil {
		t.Fatal(err.Error())
	}
	searchs := map[string]Search{
		"bleve": NewSearchBleve(index),
	}
	t.Parallel()
	for i := range searchs {
		search := searchs[i]
		t.Run(i, func(t *testing.T) {
			t.Run("insert", func(t *testing.T) {
				require.Nil(t, search.Insert(&SearchInsertItem{"foo", "foo", "test"}))
				require.Nil(t, search.Insert(&SearchInsertItem{"bar", "bar", "test"}))
				require.Nil(t, search.Insert(&SearchInsertItem{"baz", "baz", "test"}))
				require.Nil(t, search.Insert(&SearchInsertItem{"qux", "qux", "test"}))
			})

			t.Run("insert batch", func(t *testing.T) {
				items := []*SearchInsertItem{
					{"anjay", "anjay", "test"},
					{"anu", "anu", "test"},
					{"nonsi", "nonsi", "test"},
				}
				require.Nil(t, search.InsertBatch(items))
			})

			t.Run("search", func(t *testing.T) {
				ids, err := search.Query("ba")
				require.Nil(t, err)
				require.Contains(t, ids, "bar")
				require.Contains(t, ids, "baz")

				ids, err = search.Query("anjay")
				require.Nil(t, err)
				require.Contains(t, ids, "anjay")

				ids, err = search.Query("ccc")
				require.Nil(t, err)
				require.Equal(t, len(ids), 0)
			})
		})
	}
}
