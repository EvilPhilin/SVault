package db

import (
	"github.com/EvilPhilin/SVault/cfg"
	"github.com/EvilPhilin/SVault/models/vault"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func GetConnection() *gorm.DB {
	if db == nil {
		panic("No db connection!")
	}

	return db
}

func Initialize() {
	var err error
	db, err = gorm.Open(postgres.Open(cfg.ConnectionStr), &gorm.Config{})
	if err != nil {
		panic("Cant connect to db!\n" + err.Error())
	}

	err = db.AutoMigrate(&vault.Vault{})
	if err != nil {
		panic("Migration error!\n" + err.Error())
	}
}