package vault

import (
	"gorm.io/gorm"
)

type Vault struct {
	gorm.Model
	ID 					string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    AccessTokenDigest string `gorm:"type:char(128);not null"`
    Key 				string `gorm:"type:char(128);not null"`
    Path 				string `gorm:"type:varchar(150);not null"`
}