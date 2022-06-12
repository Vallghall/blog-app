package handlers

import (
	"blog-app/service"
)

type handlers struct {
	*service.Services
}

func New(s *service.Services) *handlers {
	return &handlers{s}
}
