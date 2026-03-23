package handler

import (
	"gateway/internal/service/auth"
	"gateway/internal/transport/http_gin/dto"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service auth.AuthService
}

func NewAuthHandler(s auth.AuthService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Register(c, req.Email, req.Password, req.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(201)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Login(c, req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(200, tokens)
}
