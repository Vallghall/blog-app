package repo

import (
	"blog-app/model/post"
	"blog-app/model/users"
	"github.com/jmoiron/sqlx"
	"os/user"
)

type AuthRepo interface {
	CreateUser(user users.User)
	GetUser(username, pw string) users.User
	GetUserById() user.User
}

type BlogRepo interface {
	CreatePost(p post.Post)
	UpdatePost(p post.Post)
	DeletePost(postId int)

	GetAllUserPosts(userId int) []post.Post
	GetLastNUserPosts(userId, n int) []post.Post
	GetPostById(id int) post.Post
}

type Repo struct {
	AuthRepo
	BlogRepo
}

func New(db *sqlx.DB) *Repo {
	return &Repo{}
}
