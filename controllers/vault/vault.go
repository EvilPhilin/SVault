package vault

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewVault(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"Vault's id": 			"some id",
		"Vault's access token": "access token",
	})
}