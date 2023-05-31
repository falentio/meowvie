package internal_test

import (
	"errors"
	. "meowvie/internal"
	"meowvie/internal/internalfakes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMovieService(t *testing.T) {
	movieRepo := &internalfakes.FakeMovieRepo{}
	downloadUrlRepo := &internalfakes.FakeDownloadUrlRepo{}
	search := &internalfakes.FakeSearch{}
	signer := &internalfakes.FakeSigner{}
	reset := func() {
		*movieRepo = internalfakes.FakeMovieRepo{}
		*downloadUrlRepo = internalfakes.FakeDownloadUrlRepo{}
		*search = internalfakes.FakeSearch{}
		*signer = internalfakes.FakeSigner{}
	}
	ms := NewMovieService(movieRepo, downloadUrlRepo, search, signer)
	t.Run("query", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			reset()
			search.QueryReturns([]string{}, nil)
			movies, err := ms.Query("")
			require.Nil(t, err)
			require.NotNil(t, movies)
		})

		t.Run("failed", func(t *testing.T) {
			reset()
			search.QueryReturns(nil, errors.New("testing"))
			movies, err := ms.Query("")
			require.NotNil(t, err)
			require.Nil(t, movies)
			require.EqualError(t, err, "testing")
		})
	})

	t.Run("create", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			movieRepo.CreateReturns(nil)
			downloadUrlRepo.CreateReturns(nil)
			search.InsertBatchReturns(nil)
			signer.CompareReturns(nil)
			m := &Movie{}
			result, err := ms.Create(m, "testing")
			require.Nil(t, err)
			require.NotNil(t, result)
		})
		t.Run("error", func(t *testing.T) {
			reset()
			movieRepo.CreateReturns(nil)
			downloadUrlRepo.CreateReturns(nil)
			search.InsertBatchReturns(nil)
			signer.CompareReturns(errors.New("hhh"))
			m := &Movie{}
			result, err := ms.Create(m, "testing")
			require.Nil(t, result)
			require.NotNil(t, err)
			require.EqualError(t, err, "Invalid Signature")
			require.Equal(t, 0, search.InsertCallCount())
			require.Equal(t, 0, movieRepo.CreateCallCount())
			require.Equal(t, 0, downloadUrlRepo.CreateCallCount())
		})
	})
}
