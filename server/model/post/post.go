package post

import "time"

type Post struct {
	Id       int       `json:"id" db:"id"`
	AuthorId int       `json:"author_id" db:"author_id"`
	Title    string    `json:"title" db:"title"`
	Content  string    `json:"content" db:"content"`
	Date     time.Time `json:"date" db:"date"`
	Hashtags []string  `json:"hashtags" db:"hashtags"`
}
