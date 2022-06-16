package handlers

import "github.com/gin-gonic/gin"

func (h *handlers) HandleRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", h.index)

	auth := r.Group("/auth")
	{
		auth.POST("sign-in")  // login endpoint
		auth.POST("sign-up")  // registration endpoint
		auth.POST("refresh")  // token pair refreshment
		auth.POST("sign-out") // logging out
	}

	blog := r.Group("/blog")
	{
		blog.GET(":user_id/*num") // Getting the *num posts from a user's blog
		blog.POST(":user_id")     // Adding a post to a user's blog
		blog.PUT(":post_id")      // Updating a post's content
		blog.DELETE(":post_id")   // Deleting a post
	}

	return r
}
