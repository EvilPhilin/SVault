package vault

import (
	"net/http"

	"github.com/EvilPhilin/SVault/cfg"
	"github.com/EvilPhilin/SVault/encryptor"
	db "github.com/EvilPhilin/SVault/models/repository"
	"github.com/EvilPhilin/SVault/models/vault"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateNewVault(c *gin.Context) {
	id, accessToken := uuid.New(), encryptor.CreateToken(64)
	vault := vault.Vault{
		ID: id.String(),
		Path: cfg.DataPath,
		AccessTokenDigest: encryptor.GetHash(accessToken),
		Key: encryptor.CreateToken(64),
	}

	if res := db.GetConnection().Create(&vault); res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cant create vault =(",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Vault's id": 			id,
		"Vault's access token": accessToken,
	})
}

func DeleteVault(c *gin.Context) {
	v := c.Keys["vault"].(vault.Vault)

	if res := db.GetConnection().Delete(&v); res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cant delete your vault =(",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Deleted successfully",
	})
}

func RewriteVault(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Vault's content is rewritten successfully",
		"vault": id,
	})
}

func AppendToVault(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Appended new data successfully",
		"vault": id,
	})
}

func GetVaultContent(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"data": "Your data",
		"vault": id,
	})
}