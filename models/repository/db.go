package db

import (
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
	// Oh no, my password is compromised =(
	dsn := "host=localhost user=go_connection dbname=svault password=5566 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cant connect to db!\n" + err.Error())
	}

	err = db.AutoMigrate(&vault.Vault{})
	if err != nil {
		panic("Migration error!\n" + err.Error())
	}
}