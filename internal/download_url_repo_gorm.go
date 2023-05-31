package internal

import (
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type downloadUrlRepoGorm struct {
	DB *gorm.DB
}

func NewDownloadUrlRepoGorm(DB *gorm.DB) DownloadUrlRepo {
	notNil(DB, "NewDownloadUrlRepoGorm.DB")
	return &downloadUrlRepoGorm{DB}
}

func (repo *downloadUrlRepoGorm) Create(d *DownloadUrl) error {
	d.ID = xid.New()
	return repo.DB.Create(d).Error
}

func (repo *downloadUrlRepoGorm) CreateBatch(ds []*DownloadUrl) error {
	if len(ds) == 0 {
		return nil
	}
	for _, d := range ds {
		d.ID = xid.New()
	}
	return repo.DB.Create(ds).Error
}

func (repo *downloadUrlRepoGorm) FindByMovieID(id xid.ID) ([]*DownloadUrl, error) {
	ds := make([]*DownloadUrl, 5)
	err := repo.DB.Find(&ds, "movie_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return ds, err
}

func (repo *downloadUrlRepoGorm) Find(id xid.ID) (*DownloadUrl, error) {
	d := &DownloadUrl{}
	d.ID = id
	err := repo.DB.First(d).Error
	if err != nil {
		return nil, err
	}
	return d, nil
}
