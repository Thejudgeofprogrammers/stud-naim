package handler

import (
	"gateway/internal/service/favorite"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	service favorite.FavoriteService
}

func NewFavoriteHandler(s favorite.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{service: s}
}

func (h *FavoriteHandler) List(c *gin.Context) {
	userID := c.GetString("user_id")
	opps, err := h.service.List(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, opps)
}

func (h *FavoriteHandler) Add(c *gin.Context) {
	userID := c.GetString("user_id")
	oppID := c.Param("id")
	if err := h.service.Add(c, userID, oppID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *FavoriteHandler) Remove(c *gin.Context) {
	userID := c.GetString("user_id")
	oppID := c.Param("id")
	if err := h.service.Remove(c, userID, oppID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}