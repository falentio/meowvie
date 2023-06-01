package internal

import (
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type movieRepoGorm struct {
	DB *gorm.DB
}

func NewMovieRepoGorm(DB *gorm.DB) MovieRepo {
	notNil(DB, "MovieRepoGorm.DB")
	return &movieRepoGorm{DB}
}

func (repo *movieRepoGorm) Create(m *Movie) error {
	m.ID = xid.New()
	return repo.DB.Omit("DownloadUrl").Create(m).Error
}

func (repo *movieRepoGorm) Find(id xid.ID) (*Movie, error) {
	m := &Movie{}
	m.ID = id
	err := repo.DB.Preload("DownloadUrl").First(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (repo *movieRepoGorm) FindAll(ids []xid.ID) ([]*Movie, error) {
	ms := make([]*Movie, len(ids))
	for i, id := range ids {
		ms[i] = &Movie{}
		ms[i].ID = id
		if err := repo.DB.Preload("DownloadUrl").Take(ms[i]).Error; err != nil {
			return nil, err
		}
	}
	return ms, nil
}

func (repo *movieRepoGorm) Delete(id xid.ID) error {
	m := &Movie{}
	m.ID = id
	return repo.DB.Delete(m).Error
}
