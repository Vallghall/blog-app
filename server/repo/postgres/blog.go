package postgres

import (
	"blog-app/model/post"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
)

type BlogPostgres struct {
	db *sqlx.DB
}

func NewBlogRepo(db *sqlx.DB) *BlogPostgres {
	return &BlogPostgres{db}
}

func (bp *BlogPostgres) CreatePost(p post.Post) {
	query := fmt.Sprintf(`
INSERT INTO %s (
	author_id,
	title,
	content,
	date,
	hashtags
) VALUES ($1,$2,$3,$4,$5);`, postsTable)
	_, err := bp.db.Query(query, p.AuthorId, p.Title, p.Content, p.Date, pq.Array(p.Hashtags))
	if err != nil {
		log.Println("ERROR: " + err.Error())
	}
}

func (bp *BlogPostgres) UpdatePost(p post.Post) {
	//TODO implement me
	panic("implement me")
}

func (bp *BlogPostgres) DeletePost(postId int) {
	//TODO implement me
	panic("implement me")
}

func (bp *BlogPostgres) GetAllUserPosts(userId int) []post.Post {
	//TODO implement me
	panic("implement me")
}

func (bp *BlogPostgres) GetLastNUserPosts(userId, n int) []post.Post {
	//TODO implement me
	panic("implement me")
}

func (bp *BlogPostgres) GetPostById(id int) post.Post {
	//TODO implement me
	panic("implement me")
}
