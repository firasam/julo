package handler

import (
	"net/http"

	"github.com/firasam/julo/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ViewTransaction(c *gin.Context) {
	xid := c.GetString("xid")

	if xid == "" {
		c.JSON(http.StatusUnauthorized, dto.FormatFail("Not authorized"))
		return
	}

	transaction, err := h.transactionService.ViewTransaction(xid)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.FormatFail(err.Error()))
		return
	}
	c.JSON(http.StatusAccepted, dto.FormatViewTransactionResponse(transaction))
}

func (h *Handler) Deposit(c *gin.Context) {
	xid := c.GetString("xid")
	input := &dto.DepositRequest{}

	err := c.ShouldBind(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.FormatFail(err.Error()))
		return
	}
	print(input.ReferenceID)
	transaction, err := h.transactionService.Deposit(xid, input)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.FormatFail(err.Error()))
		return
	}
	c.JSON(http.StatusAccepted, dto.FormatDepositResponse(transaction))
}

func (h *Handler) Withdraw(c *gin.Context) {
	xid := c.GetString("xid")
	input := &dto.WithdrawRequest{}

	err := c.ShouldBind(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.FormatFail(err.Error()))
		return
	}

	transaction, err := h.transactionService.Withdraw(xid, input)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.FormatFail(err.Error()))
		return
	}
	c.JSON(http.StatusAccepted, dto.FormatDepositResponse(transaction))
}
