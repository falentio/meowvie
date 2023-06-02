package internal

import (
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/rs/xid"
)

type movieRepoLru struct {
	cache *lru.ARCCache[string, *Movie]
	repo  MovieRepo
}

func NewMovieRepoLru(repo MovieRepo) MovieRepo {
	cache, _ := lru.NewARC[string, *Movie](3000)
	return &movieRepoLru{cache, repo}
}

func (r *movieRepoLru) Find(id xid.ID) (*Movie, error) {
	if r.cache.Contains(id.String()) {
		m, _ := r.cache.Get(id.String())
		return m, nil
	}
	m, err := r.repo.Find(id)
	if err != nil {
		return nil, err
	}
	r.cache.Add(id.String(), m)
	return m, nil
}

func (r *movieRepoLru) FindAll(ids []xid.ID) ([]*Movie, error) {
	ms := make([]*Movie, len(ids))
	for i := range ids {
		m, err := r.Find(ids[i])
		if err != nil {
			return nil, err
		}
		ms[i] = m
	}
	return ms, nil
}

func (r *movieRepoLru) Delete(id xid.ID) error {
	r.cache.Remove(id.String())
	return r.repo.Delete(id)
}

func (r *movieRepoLru) Create(m *Movie) error {
	return r.repo.Create(m)
}
