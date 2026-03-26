package handler

import (
	"gateway/internal/domain"
	"gateway/internal/service/opportunity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OpportunityHandler struct {
	service opportunity.OpportunityService
}

func NewOpportunityHandler(s opportunity.OpportunityService) *OpportunityHandler {
	return &OpportunityHandler{
		service: s,
	}
}

func (h *OpportunityHandler) Create(c *gin.Context) {
	var req domain.Opportunity

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("user_id")
	role := domain.Role(c.GetString("role"))

	req.CompanyID = userID

	if err := h.service.Create(c, &req, role); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *OpportunityHandler) Get(c *gin.Context) {
	id := c.Param("id")

	o, err := h.service.Get(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, o)
}

func (h *OpportunityHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var req domain.Opportunity
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("user_id")

	req.ID = id

	err := h.service.Update(c, &req, userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *OpportunityHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("user_id")

	err := h.service.Delete(c, id, userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *OpportunityHandler) List(c *gin.Context) {
	list, err := h.service.List(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *OpportunityHandler) Filter(c *gin.Context) {
	tag := c.Query("tag")
	format := domain.WorkFormat(c.Query("format"))

	list, err := h.service.Filter(c, tag, format)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}
