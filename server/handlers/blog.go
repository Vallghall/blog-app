package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handlers) createPost(c *gin.Context) {

}

func (h *handlers) readPosts(c *gin.Context) {
	id, num := c.GetInt("user_id"), c.GetInt("num")
	c.JSON(http.StatusOK, map[string]int{
		"id":  id,
		"num": num,
	})
}

func (h *handlers) updatePost(c *gin.Context) {

}

func (h *handlers) deletePost(c *gin.Context) {

}
