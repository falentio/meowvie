package internal

import "github.com/rs/xid"

type DownloadUrl struct {
	Model

	Url        string `json:"url"`
	Server     string `json:"server"`
	Resolution string `json:"resolution"`

	MovieID xid.ID `json:"movieId" gorm:"index"`
	Movie   *Movie `json:"movie"`
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DownloadUrlRepo

type DownloadUrlRepo interface {
	Create(d *DownloadUrl) error
	CreateBatch(ds []*DownloadUrl) error
	FindByMovieID(id xid.ID) ([]*DownloadUrl, error)
	Find(id xid.ID) (*DownloadUrl, error)
}
