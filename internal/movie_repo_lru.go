package internal

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

type movieRepoLru struct {
	cache lru.ARCCache[string, *Movie]
	repo  MovieRepo
}

// func NewMovieRepoLru(repo MovieRepo) MovieRepo {
// 	return &movieRepoLru{lru.NewARC[string, *Movie](10), repo}
// }
