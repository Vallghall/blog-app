package handlers

import (
	"blog-app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrInvalidParams = "error: invalid params"
)

type handlers struct {
	*service.Services
}

func New(s *service.Services) *handlers {
	return &handlers{s}
}

func (h *handlers) index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "Hello. This is a server of a blog application. Check the link for the API reference",
		"link":    "github.com/Vallghall/blog-app",
	})
}
