package internal_test

import (
	. "meowvie/internal"
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/require"
)

func TestDownloadUrlRepo(t *testing.T) {
	repos := map[string]DownloadUrlRepo{
		"gorm": NewDownloadUrlRepoGorm(db),
	}
	t.Parallel()
	for i := range repos {
		repo := repos[i]
		t.Run(i, func(t *testing.T) {
			d := &DownloadUrl{
				Url:        "https://example.com",
				Server:     "foo",
				Resolution: "720p",
				MovieID:    xid.New(),
			}

			t.Run("create", func(t *testing.T) {
				err := repo.Create(d)
				require.Nil(t, err)
			})

			t.Run("create batch", func(t *testing.T) {
				err := repo.CreateBatch([]*DownloadUrl{d})
				require.Nil(t, err)
				err = repo.CreateBatch([]*DownloadUrl{})
				require.Nil(t, err)
			})

			t.Run("find by movie id", func(t *testing.T) {
				ds, err := repo.FindByMovieID(d.MovieID)
				require.Nil(t, err)
				require.Equal(t, 2, len(ds))

				ds, err = repo.FindByMovieID(xid.New())
				require.Nil(t, err)
				require.Equal(t, 0, len(ds))
			})

			t.Run("find", func(t *testing.T) {
				stored, err := repo.Find(d.ID)
				require.Nil(t, err)
				require.Equal(t, d.MovieID, stored.MovieID)

				_, err = repo.Find(xid.New())
				require.NotNil(t, err)
			})
		})
	}
}
