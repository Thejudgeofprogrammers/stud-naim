package handler

import (
	"gateway/internal/service/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseHandler struct {
	service response.ResponseService
}

func NewResponseHandler(s response.ResponseService) *ResponseHandler {
	return &ResponseHandler{service: s}
}

type CreateResponseRequest struct {
	OpportunityID string `json:"opportunity_id"`
	Message       string `json:"message,omitempty"`
}

func (h *ResponseHandler) Create(c *gin.Context) {
	userID := c.GetString("user_id")
	var req CreateResponseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(c, userID, req.OpportunityID, req.Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *ResponseHandler) List(c *gin.Context) {
	userID := c.GetString("user_id")
	resps, err := h.service.ListByUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resps)
}