package internal_test

import (
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/require"

	. "meowvie/internal"
)

func TestMovieRepo(t *testing.T) {
	repos := map[string]MovieRepo{
		"gorm": NewMovieRepoGorm(db),
	}

	t.Parallel()
	for name := range repos {
		repo := repos[name]
		t.Run(name, func(t *testing.T) {
			m := &Movie{
				Provider:     "test",
				Title:        "testing",
				PageUrl:      "https://example.com",
				ThumbnailUrl: "https://example.com",
			}
			t.Run("create", func(t *testing.T) {
				err := repo.Create(m)
				require.Nil(t, err)
				require.NotZero(t, m.ID)
			})

			t.Run("providerList", func(t *testing.T) {
				providers, err := repo.ProviderList()
				require.Nil(t, err)
				require.Contains(t, providers, "test")
			})

			t.Run("find", func(t *testing.T) {
				stored, err := repo.Find(m.ID)
				require.Nil(t, err)
				require.NotNil(t, stored)
				require.Equal(t, m.ID, stored.ID)
				require.Equal(t, m.Title, stored.Title)

				stored, err = repo.Find(xid.New())
				require.NotNil(t, err)
				require.Nil(t, stored)
			})

			t.Run("find all", func(t *testing.T) {
				stored, err := repo.FindAll([]xid.ID{m.ID})
				require.Nil(t, err)
				require.Equal(t, 1, len(stored))
				require.Equal(t, m.ID, stored[0].ID)
			})

			t.Run("delete", func(t *testing.T) {
				err := repo.Delete(m.ID)
				require.Nil(t, err)

				_, err = repo.Find(m.ID)
				require.NotNil(t, err)
			})
		})
	}
}
