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
	//TODO implement me
	panic("implement me")
}

func (b *Blog) UpdatePost(p post.Post) {
	//TODO implement me
	panic("implement me")
}

func (b *Blog) DeletePost(postId int) {
	//TODO implement me
	panic("implement me")
}

func (b *Blog) GetAllUserPosts(userId int) []post.Post {
	//TODO implement me
	panic("implement me")
}

func (b *Blog) GetLastNUserPosts(userId, n int) []post.Post {
	//TODO implement me
	panic("implement me")
}

func (b *Blog) GetPostById(id int) post.Post {
	//TODO implement me
	panic("implement me")
}
