package handlers

import "github.com/gin-gonic/gin"

func (h *handlers) HandleRoutes() *gin.Engine {
	r := gin.Default()
	return r
}
