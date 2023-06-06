package internal

import "github.com/rs/xid"

type DownloadUrl struct {
	Model

	Url        string `json:"url"`
	Server     string `json:"server"`
	Resolution string `json:"resolution"`
	Size       string `json:"size"`

	MovieID xid.ID `json:"movieId" gorm:"index"`
	Movie   *Movie `json:"movie"`
}

//counterfeiter:generate . DownloadUrlRepo

type DownloadUrlRepo interface {
	Create(d *DownloadUrl) error
	CreateBatch(ds []*DownloadUrl) error
	FindByMovieID(id xid.ID) ([]*DownloadUrl, error)
	Find(id xid.ID) (*DownloadUrl, error)
}
