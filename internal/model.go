package internal

import (
	"time"

	"github.com/rs/xid"

	"gorm.io/gorm"
)

type Model struct {
	ID        xid.ID         `gorm:"primaryKey,type:varchar(20)" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
