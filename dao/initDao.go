package dao

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"articleManager/conf"
	"time"
)

var(
	db *gorm.DB
	redisDb *redis.ClusterClient
)

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


//初始化redis
func InitRedis(conf *conf.Conf){
	redisopts:=&redis.ClusterOptions{
		Addrs:              conf.RedisAddr,
		Password:           conf.Password,
		PoolSize:           100,
		MinIdleConns:       16,
		IdleTimeout:        240*time.Second,
	}
	redisDb=redis.NewClusterClient(redisopts)
}