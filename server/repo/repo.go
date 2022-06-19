package repo

import (
	"blog-app/model/post"
	"blog-app/model/users"
	"blog-app/repo/postgres"
	"github.com/jmoiron/sqlx"
)

type AuthRepo interface {
	CreateUser(user users.User) int
	GetUser(username, pw string) users.User
	GetUserById(id int) users.User
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
	return &Repo{
		postgres.NewAuthRepo(db),
		postgres.NewBlogRepo(db),
	}
}
