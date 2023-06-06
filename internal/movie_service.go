package internal

import (
	"errors"

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
	for _, d := range m.DownloadUrl {
		d.MovieID = m.ID
	}
	if err := ms.DownloadUrlRepo.CreateBatch(m.DownloadUrl); err != nil {
		return nil, err
	}
	if err := ms.Search.Insert(&SearchInsertItem{m.ID.String(), m.Title, m.Provider}); err != nil {
		return nil, err
	}
	return m, nil
}

func (ms *MovieService) Query(term string) ([]*Movie, error) {
	ids, err := ms.Search.Query(term)
	if err != nil {
		return nil, err
	}
	xids, err := toXid(ids)
	if err != nil {
		return nil, errors.New("bad id stored in database")
	}
	return ms.MovieRepo.FindAll(xids)
}

func (ms *MovieService) Find(id xid.ID) (*Movie, error) {
	return ms.MovieRepo.Find(id)
}

func (ms *MovieService) Delete(id xid.ID, signature string) error {
	if err := ms.Signer.Compare(id.String(), signature); err != nil {
		return fiber.NewError(400, "Invalid Signature")
	}
	return ms.MovieRepo.Delete(id)
}
