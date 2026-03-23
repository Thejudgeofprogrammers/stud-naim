package app

import (
	"gateway/internal/config"

	"github.com/gin-gonic/gin"
)

func NewApp(env *config.Config) *gin.Engine {
	r := gin.New()
	
	RegisterRoutes(r, env)

	return r
}