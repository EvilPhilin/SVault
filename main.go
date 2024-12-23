package main

import (
	"github.com/EvilPhilin/SVault/controllers/vault"
	"github.com/EvilPhilin/SVault/middleware"
	db "github.com/EvilPhilin/SVault/models/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Initialize()

	router := gin.Default()

	router.POST("/vaults", middleware.CheckSignet, vault.CreateNewVault)
	vaultGroup := router.Group("/vaults/:id", middleware.CheckAuthToken)
	{
		vaultGroup.DELETE("/", vault.DeleteVault)
		vaultGroup.PUT("/", vault.RewriteVault)
		vaultGroup.PATCH("/", vault.AppendToVault)
		vaultGroup.GET("/", vault.GetVaultContent)
	}

	router.Run()
}
