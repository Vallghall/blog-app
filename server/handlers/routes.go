package handlers

import "github.com/gin-gonic/gin"

func (h *handlers) HandleRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", h.index)

	auth := r.Group("/auth")
	{
		auth.POST("sign-in", h.signIn)   // login endpoint
		auth.POST("sign-up", h.signUp)   // registration endpoint
		auth.POST("refresh", h.refresh)  // token pair refreshment
		auth.POST("sign-out", h.signOut) // logging out
	}

	blog := r.Group("/blog")
	{
		blog.GET(":user_id/:num", h.readPosts) // Getting the *num posts from a user's blog
		blog.POST(":user_id", h.createPost)    // Adding a post to a user's blog
		blog.PUT(":post_id", h.updatePost)     // Updating a post's content
		blog.DELETE(":post_id", h.deletePost)  // Deleting a post
	}

	return r
}
