package handlers

import (
	"blog-app/model/post"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (h *handlers) createPost(c *gin.Context) {
	uid, ok := c.Get(UID)
	id, err := strconv.Atoi(c.Param("user_id"))
	if uid != id || !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthorized)
		return
	}

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrInvalidParams)
		return
	}

	var p post.Post
	err = c.BindJSON(&p)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrInvalidParams)
		return
	}
	p.AuthorId = id
	p.Date = time.Now().UTC()

	h.BlogService.CreatePost(p)
	c.JSON(http.StatusOK, map[string]string{
		"message": "OK",
	})
}

func (h *handlers) readPosts(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrInvalidParams)
		return
	}

	num, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrInvalidParams)
		return
	}

	posts := h.BlogService.GetLastNUserPosts(id, num)
	c.JSON(http.StatusOK, map[string][]post.Post{
		"posts": posts,
	})
}

func (h *handlers) getPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrInvalidParams)
		return
	}

	p := h.BlogService.GetPostById(id)
	c.JSON(http.StatusOK, map[string]post.Post{
		"post": p,
	})
}

func (h *handlers) updatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrInvalidParams)
		return
	}
	usrId, _ := c.Get(UID)
	prvVersion := h.BlogService.GetPostById(id)
	if usrId != prvVersion.AuthorId {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"message": "Cannot edit another user's post",
		})
		return
	}

	var p post.Post
	err = c.BindJSON(&p)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrInvalidParams)
		return
	}
	p.Id = id

	h.BlogService.UpdatePost(p)
	c.JSON(http.StatusOK, map[string]string{
		"message": "OK",
	})
}

func (h *handlers) deletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrInvalidParams)
		return
	}

	usrId, _ := c.Get(UID)
	prvVersion := h.BlogService.GetPostById(id)
	if usrId != prvVersion.AuthorId {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"message": "Cannot delete another user's post",
		})
		return
	}

	h.BlogService.DeletePost(id)
	c.JSON(http.StatusOK, map[string]string{
		"message": "OK",
	})
}
