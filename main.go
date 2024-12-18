package main

import (
	"github.com/EvilPhilin/SVault/controllers/vault"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	vaultGroup := router.Group("/vaults")
	{
		vaultGroup.POST("/", vault.CreateNewVault)
	}

	router.Run()
}
