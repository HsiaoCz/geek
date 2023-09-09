package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Auther          string `gorm:"columnL:auther;type:varchar(20);" json:"auther"`
	ArticleIdentity int    `gorm:"column:article_identity;type:int(11);" json:"article_identity"`
	UserIdentity    int    `gorm:"column:user_identity;type:int(11);" json:"user_identity"`
	Title           string `gorm:"column:title;type:varchar(50);" json:"title"`
	ArticleType     string `gorm:"column:article_type;type:varchar(20);" json:"article_type"`
	Content         string `gorm:"column:content;type:varchar(50000);" json:"content"`
	Likes           int    `gorm:"column:likes;type:int(11);" json:"likes"`
}

func (a Article) TableName() string {
	return "article"
}
