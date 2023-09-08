package model

import "time"

type Article struct {
	Auther          string    `json:"auther"`
	ArticleIdemtity int       `json:"article_identity"`
	UserIdentity    int       `json:"user_identity"`
	Title           string    `json:"title"`
	ArticleType     string    `json:"article_type"`
	Content         string    `json:"content"`
	CreateTime      time.Time `json:"create_time"`
	Likes           int       `json:"likes"`
}
