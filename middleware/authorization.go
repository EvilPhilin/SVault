package middleware

import (
	"net/http"

	"github.com/EvilPhilin/SVault/encryptor"
	db "github.com/EvilPhilin/SVault/models/repository"
	"github.com/EvilPhilin/SVault/models/vault"
	"github.com/gin-gonic/gin"
)

/*
	Checks whether provided token allows creation of new vault,
	for now there's no real logic in here
*/
func CheckSignet(c *gin.Context) {
	if signet := c.PostForm("signet"); len(signet) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect signet",
		})
		c.Abort()
	}
}

func CheckAuthToken(c *gin.Context) {
	vaultId, authToken := c.Param("id"), c.Query("authorization_token")

	if len(vaultId) != 36 || len(authToken) != 128 {
		returnIncorrectParams(c)
	}

	var v vault.Vault
	tokenDigest := encryptor.GetHash(authToken)
	db.GetConnection().
		Where("id = ? AND access_token_digest = ?", vaultId, tokenDigest).
		First(&v)

	if v.Path == "" {
		returnIncorrectParams(c)
	}

	c.Set("vault", v)
}

func returnIncorrectParams(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Incorrect vault or token",
	})
	c.Abort()
}