package handler

import (
	"net/http"

	"github.com/firasam/julo/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	input := &dto.WalletCreateRequest{}

	err := c.ShouldBind(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.FormatFail(err.Error()))
		return
	}

	_, err = h.walletService.CreateWallet(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.FormatFail(err.Error()))
		return
	}
	token, err := h.jwtService.GenerateToken(input.CustomerXID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.FormatFail(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dto.FormatCreatedDetail(token))
}

func (h *Handler) EnableWallet(c *gin.Context) {
	xid := c.GetString("xid")

	if xid == "" {
		c.JSON(http.StatusUnprocessableEntity, dto.FormatFail("xid not set"))
		return
	}

	wallet, err := h.walletService.EnableWallet(xid)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.FormatFail(err.Error()))
		return
	}
	c.JSON(http.StatusAccepted, dto.FormatedEnable(wallet))
}

func (h *Handler) DisableWallet(c *gin.Context) {
	xid := c.GetString("xid")

	if xid == "" {
		c.JSON(http.StatusUnauthorized, dto.FormatFail("Not authorized"))
		return
	}

	wallet, err := h.walletService.DisableWallet(xid)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.FormatFail(err.Error()))
		return
	}
	c.JSON(http.StatusAccepted, dto.FormatedDisable(wallet))
}

func (h *Handler) ViewWallet(c *gin.Context) {
	xid := c.GetString("xid")

	if xid == "" {
		c.JSON(http.StatusUnauthorized, dto.FormatFail("Not authorized"))
		return
	}

	wallet, err := h.walletService.ViewWallet(xid)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.FormatFail(err.Error()))
		return
	}
	c.JSON(http.StatusAccepted, dto.FormatedEnable(wallet))
}
