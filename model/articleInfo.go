package model

import "time"

//文章信息
type ArticleInfo struct {
	Id              int32     `gorm:"column:id;primary_key"`
	Article_context string    `gorm:"column:article"`
	Type            int32     `gorm:"column:type"`
	Photo           string    `gorm:"column:photo"`
	Insert_time     time.Time `gorm:"column:insert_time"`
	Update_time     time.Time `gorm:"column:update_time"`
}

func (ArticleInfo) TableName() string {
	return "article"
}