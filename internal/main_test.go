package internal_test

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	. "meowvie/internal"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	db, err = gorm.Open(sqlite.Open("file:../database/testing.db"), &gorm.Config{})
	if err != nil {
		panic("failed to open database " + err.Error())
	}
	if err := db.AutoMigrate(&Movie{}, &DownloadUrl{}); err != nil {
		panic("failed to run migration " + err.Error())
	}
	m.Run()
}
