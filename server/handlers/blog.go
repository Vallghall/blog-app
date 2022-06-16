package handlers

import (
	"blog-app/model/post"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *handlers) createPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var p post.Post
	err = c.BindJSON(&p)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	p.AuthorId = id

	h.BlogService.CreatePost(p)
	c.JSON(http.StatusOK, map[string]string{
		"message": "OK",
	})
}

func (h *handlers) readPosts(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	num, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	posts := h.BlogService.GetLastNUserPosts(id, num)
	c.JSON(http.StatusOK, map[string][]post.Post{
		"posts": posts,
	})
}

func (h *handlers) updatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var p post.Post
	err = c.BindJSON(&p)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	p.AuthorId = id

	h.BlogService.UpdatePost(p)
	c.JSON(http.StatusOK, map[string]string{
		"message": "OK",
	})
}

func (h *handlers) deletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	h.BlogService.DeletePost(id)
	c.JSON(http.StatusOK, map[string]string{
		"message": "OK",
	})
}
