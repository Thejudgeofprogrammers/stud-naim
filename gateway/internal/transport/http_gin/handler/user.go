package handler

import (
	"gateway/internal/service/user"
	"gateway/internal/transport/http_gin/dto"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(s user.UserService) *UserHandler {
	return &UserHandler{
		userService: s,
	}
}

func (h *UserHandler) ListStudents(c *gin.Context) {

}

func (h *UserHandler) GetStudent(c *gin.Context) {
	id := c.Param("id")

	student, err := h.userService.GetStudent(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, student)
}

func (h *UserHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.UpdateStudent(c, id, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *UserHandler) ListEmployers(c *gin.Context) {

}

func (h *UserHandler) GetEmployer(c *gin.Context) {
	id := c.Param("id")

	student, err := h.userService.GetEmployer(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, student)
}

func (h *UserHandler) UpdateEmployer(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateEmployerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.UpdateEmployer(c, id, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *UserHandler) GetMe(c *gin.Context) {
	userID := c.GetString("user_id")

	c.JSON(200, gin.H{
		"id": userID,
	})
}
