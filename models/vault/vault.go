package vault

import (
	"os"
	"path/filepath"

	"github.com/EvilPhilin/SVault/cfg"
	"gorm.io/gorm"
)

type Vault struct {
	gorm.Model
	ID 					string `gorm:"type:uuid;primaryKey"`
    AccessTokenDigest 	string `gorm:"type:char(128);not null"`
    Key 				string `gorm:"type:char(128);not null"`
    Path 				string `gorm:"type:varchar(150);not null"`
}

func (v *Vault) CreateFile() (error) {
	fullPath := filepath.Join(cfg.DataPath, v.ID)

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}