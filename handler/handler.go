package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kjasuquo/preparation/database"
	"github.com/kjasuquo/preparation/model"
	"net/http"
	"strconv"
)

type handler struct {
	DB database.Database
}

func NewHandler(db database.Database) *handler {
	return &handler{DB: db}
}

func (h *handler) CreateWalletHandler(c *gin.Context) {
	var wallet model.Wallet
	wallet.CustomerName = c.Query("name")

	wallet.Balance = 0
	err := h.DB.CreateWallet(wallet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot create wallet"})
		return
	}

	c.JSON(http.StatusOK, "successful")

}

func (h *handler) GetWalletHandler(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid wallet"})
		return
	}

	wallet, err := h.DB.GetWallet(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot get Wallet"})
		return
	}

	c.JSON(http.StatusOK, wallet)

}

func (h *handler) WalletTransactionHandler(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid wallet"})
		return
	}

	wallet, err := h.DB.GetWallet(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot get Wallet"})
		return
	}

	var AmountPaid model.AmountPaid

	err = c.BindJSON(&AmountPaid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	err = h.DB.WalletTransaction(wallet, AmountPaid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Insufficient funds"})
		return
	}

	c.JSON(http.StatusOK, "successful")

}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
