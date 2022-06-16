package handlers

import (
	"blog-app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handlers struct {
	*service.Services
}

func (h *handlers) index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "Hello",
	})
}

func New(s *service.Services) *handlers {
	return &handlers{s}
}
