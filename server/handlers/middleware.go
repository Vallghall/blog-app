package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func (h *handlers) authenticate(c *gin.Context) {
	header := c.GetHeader("Authorization")
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrUnauthorized+"1")
		return
	}
	at := headerParts[1]

	rt, err := c.Cookie("refresh-token")
	if err != nil || at == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthorized+"2")
		return
	}
	atUID, err := h.AuthService.ParseToken(at)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthorized+"3")
		return
	}

	rtUID, err := h.AuthService.ParseToken(rt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthorized+"4")
		return
	}

	if atUID != rtUID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthorized+"5")
		return
	}

	u := h.AuthService.GetUserById(atUID)
	if u.Name == "" {
		log.Println(atUID)
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthorized+"6")
		return
	}

	c.Set(UID, u.Id)
	c.Set(UserName, u.Nickname)
}
