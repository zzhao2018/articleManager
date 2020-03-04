package model

//类型信息
type TypeInfo struct {
	Id       int32  `gorm:"column:id;primary_key"`
	TypeName string `gorm:"column:type_name"`
	SendTime string `gorm:"send_time"`
	Delete_code int `gorm:"delete_code"`
}

func(TypeInfo) TableName() string{
	return "type"
}
