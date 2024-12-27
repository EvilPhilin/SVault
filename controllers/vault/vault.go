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

	if err := vault.CreateFile(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cant create file system entry",
		})
		c.Abort()
	}

	if res := db.GetConnection().Create(&vault); res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cant create db entry",
		})
		c.Abort()
	}

	c.JSON(http.StatusCreated, gin.H{
		"Vault's id": 			id,
		"Vault's access token": accessToken,
	})
}

func DeleteVault(c *gin.Context) {
	vId, _ := c.Get("vaultId")
	aT, _ := c.Get("authToken")
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Deleted successfully",
		"vault": vId,
		"token": aT,
	})
}

func RewriteVault(c *gin.Context) {
	vId, _ := c.Get("vaultId")
	aT, _ := c.Get("authToken")
	c.JSON(http.StatusOK, gin.H{
		"message": "Vault's content is rewritten successfully",
		"vault": vId,
		"token": aT,
	})
}

func AppendToVault(c *gin.Context) {
	vId, _ := c.Get("vaultId")
	aT, _ := c.Get("authToken")
	c.JSON(http.StatusOK, gin.H{
		"message": "Appended new data successfully",
		"vault": vId,
		"token": aT,
	})
}

func GetVaultContent(c *gin.Context) {
	vId, _ := c.Get("vaultId")
	aT, _ := c.Get("authToken")
	c.JSON(http.StatusOK, gin.H{
		"data": "Your data",
		"vault": vId,
		"token": aT,
	})
}