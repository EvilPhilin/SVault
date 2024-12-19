package middleware

import (
	"net/http"

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

	// tokenDigest = hash(token)
	// Verify that vault with this digest exists or abort
	if vaultId == "42" || len(authToken) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect vault or token",
		})
		c.Abort()
	}

	c.Set("vaultId", vaultId)
	c.Set("authToken", authToken)
}