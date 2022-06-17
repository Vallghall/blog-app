package service

import (
	"blog-app/model/post"
	"blog-app/model/users"
	"blog-app/repo"
)

type AuthService interface {
	CreateUser(user users.User)
	GetUser(username, pw string) users.User
	GetUserById(id int) users.User
}

type BlogService interface {
	CreatePost(p post.Post)
	UpdatePost(p post.Post)
	DeletePost(postId int)

	GetAllUserPosts(userId int) []post.Post
	GetLastNUserPosts(userId, n int) []post.Post
	GetPostById(id int) post.Post
}

type Services struct {
	AuthService
	BlogService
}

func New(r *repo.Repo) *Services {
	return &Services{
		AuthService: NewAuthService(r),
		BlogService: NewBlogService(r),
	}
}
