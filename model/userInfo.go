package model

type UserInfo struct {
	Id int 	`gorm:"column:id;primary_key"`
	Password string `gorm:"column:password"`
}

func(UserInfo) TableName()string{
	return "user"
}
