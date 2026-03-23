package handler

import (
	"gateway/internal/service/resume"

	"github.com/gin-gonic/gin"
)

type ResumeHandler struct {
	service resume.ResumeService
}

func NewResumeHandler(s resume.ResumeService) *ResumeHandler {
	return &ResumeHandler{
		service: s,
	}
}

func (h *ResumeHandler) GetResume(c *gin.Context) {
	id := c.Param("id")

	url, err := h.service.GetResume(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"url": url})
}

func (h *ResumeHandler) UploadResume(c *gin.Context) {
	id := c.Param("id")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file required"})
		return
	}

	err = h.service.UploadResume(c, id, file.Filename)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *ResumeHandler) DeleteResume(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteResume(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}
