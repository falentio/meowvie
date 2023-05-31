package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
)

type MovieService struct {
	MovieRepo       MovieRepo
	DownloadUrlRepo DownloadUrlRepo
	Search          Search
	Signer          Signer
}

func NewMovieService(
	m MovieRepo,
	d DownloadUrlRepo,
	s Search,
	signer Signer,
) *MovieService {
	notNil(m, "NewMovieService.MovieRepo")
	notNil(d, "NewMovieService.DownloadUrlRepo")
	notNil(s, "NewMovieService.Search")
	notNil(signer, "NewMovieService.Signer")
	return &MovieService{m, d, s, signer}
}

func (ms *MovieService) Create(m *Movie, signature string) (*Movie, error) {
	if err := ms.Signer.Compare(m.Title, signature); err != nil {
		return nil, fiber.NewError(400, "Invalid Signature")
	}
	if err := ms.MovieRepo.Create(m); err != nil {
		return nil, err
	}
	if err := ms.DownloadUrlRepo.CreateBatch(m.DownloadUrl); err != nil {
		return nil, err
	}
	if err := ms.Search.Insert(m.ID.String(), m.Title); err != nil {
		return nil, err
	}
	return m, nil
}

func (ms *MovieService) Query(term string) ([]*Movie, error) {
	ids, err := ms.Search.Query(term)
	if err != nil {
		return nil, err
	}
	movies := make([]*Movie, len(ids))
	for i := range ids {
		id, err := xid.FromString(ids[i])
		if err != nil {
			return nil, err
		}
		movies[i], err = ms.MovieRepo.Find(id)
		if err != nil {
			return nil, err
		}
	}
	return movies, nil
}

func (ms *MovieService) Find(id xid.ID) (*Movie, error) {
	return ms.MovieRepo.Find(id)
}
