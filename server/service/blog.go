package service

import (
	"blog-app/model/post"
	"blog-app/repo"
)

type Blog struct {
	r *repo.Repo
}

func NewBlogService(r *repo.Repo) *Blog {
	return &Blog{r}
}

func (b *Blog) CreatePost(p post.Post) {
	b.r.BlogRepo.CreatePost(p)
}

func (b *Blog) UpdatePost(p post.Post) {
	b.r.BlogRepo.UpdatePost(p)
}

func (b *Blog) DeletePost(postId int) {
	b.r.BlogRepo.DeletePost(postId)
}

func (b *Blog) GetAllUserPosts(userId int) []post.Post {
	return b.r.BlogRepo.GetAllUserPosts(userId)
}

func (b *Blog) GetLastNUserPosts(userId, n int) []post.Post {
	return b.r.BlogRepo.GetLastNUserPosts(userId, n)
}

func (b *Blog) GetPostById(id int) post.Post {
	return b.r.BlogRepo.GetPostById(id)
}
