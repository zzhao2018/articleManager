package dao

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"articleManager/conf"
)

var db *gorm.DB

//初始化数据库
func InitDataBase(conf *conf.Conf)error{
	urls:=conf.GetDataBaseUrl()
	var err error
	db,err=gorm.Open("mysql",urls)
	if err!=nil {
		log.Printf("InitDataBase gorm open db error,err:%+v\n",err)
		return err
	}
	return nil
}
