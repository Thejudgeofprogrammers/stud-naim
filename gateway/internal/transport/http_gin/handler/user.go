package handler

import (
	"gateway/internal/service/user"
	"gateway/internal/transport/http_gin/dto"
	"net/http"

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
	students, err := h.userService.ListStudents(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

func (h *UserHandler) GetStudent(c *gin.Context) {
	id := c.Param("id")

	student, err := h.userService.GetStudent(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *UserHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.UpdateStudent(c, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandler) ListEmployers(c *gin.Context) {
	employers, err := h.userService.ListEmployers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employers)
}

func (h *UserHandler) GetEmployer(c *gin.Context) {
	id := c.Param("id")

	employer, err := h.userService.GetEmployer(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employer)
}

func (h *UserHandler) UpdateEmployer(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateEmployerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.UpdateEmployer(c, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandler) GetMe(c *gin.Context) {
	userID := c.GetString("user_id")
	role := c.GetString("role")

	c.JSON(http.StatusOK, gin.H{
		"id":   userID,
		"role": role,
	})
}
