package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	AuthDB *gorm.DB
	LinksDB *gorm.DB
	IsConnected bool
	IsSQLite bool
	err error
)

func Initialize() *gorm.DB {
	_ = os.Mkdir("db", os.ModePerm)
	LinksDB, err = gorm.Open(sqlite.Open("db/links.db"), &gorm.Config{})
	AuthDB, err = gorm.Open(sqlite.Open("db/auth.db"), &gorm.Config{})
	
	if err != nil {
		return nil
	} else {
		IsConnected = true
		IsSQLite = true
	}

	return LinksDB
}
