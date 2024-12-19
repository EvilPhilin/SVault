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