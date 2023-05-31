package internal

import "github.com/rs/xid"

type Movie struct {
	Model

	Title        string `json:"title"`
	PageUrl      string `json:"pageUrl"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	Provider     string `json:"provider"`

	DownloadUrl []*DownloadUrl `json:"downloadUrl"`
}

type MovieSignature struct {
	*Movie
	Signature string `json:"signature"`
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . MovieRepo

type MovieRepo interface {
	Create(*Movie) error
	Find(xid.ID) (*Movie, error)
	FindAll([]xid.ID) ([]*Movie, error)
	Delete(xid.ID) error
}
