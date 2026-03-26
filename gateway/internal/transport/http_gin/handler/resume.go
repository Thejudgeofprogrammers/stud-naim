package handler

import (
	"fmt"
	"gateway/internal/service/resume"
	"net/http"
	"os"
	"path/filepath"

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
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *ResumeHandler) UploadResume(c *gin.Context) {
	id := c.Param("id")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file required"})
		return
	}

	// создаём папку если нет
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	filePath := filepath.Join(uploadDir, fmt.Sprintf("%s_%s", id, file.Filename))

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileURL := "/uploads/" + filepath.Base(filePath)

	err = h.service.UploadResume(c, id, fileURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": fileURL,
	})
}

func (h *ResumeHandler) DeleteResume(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteResume(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
