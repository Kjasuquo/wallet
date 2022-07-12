package config

import (
	"github.com/gin-gonic/gin"
	"github.com/kjasuquo/preparation/database"
	"github.com/kjasuquo/preparation/handler"
	"net/http"
)

func (cfg *config) routes() http.Handler {
	router := gin.Default()
	d := database.NewDB()
	d.SetupDB("wallet")

	h := handler.NewHandler(d)
	apirouter := router.Group("/api/v1")

	apirouter.GET("/ping", handler.PingHandler)

	apirouter.POST("/create", h.CreateWalletHandler)
	apirouter.GET("/wallet/:id", h.GetWalletHandler)
	apirouter.PUT("/transaction/:id", h.WalletTransactionHandler)

	return router
}
