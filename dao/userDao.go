package dao

import (
	"articleManager/model"
	"log"
)

//获取密码
func GetPassWord()*model.UserInfo{
	var userInfo model.UserInfo
	db.First(&userInfo)
	return &userInfo
}

//将cookiet值放入redis
func SetValueToRedis(key string,value string,timeLong int)error{
	cmd:=redisDb.Do("set",key,value,"ex",timeLong)
	_,err:=cmd.Result()
	if err!=nil {
		log.Printf("SetValueToRedis error,err:%+v\n",err)
		return err
	}
	return nil
}

//从redis中获得值
func GetValueFromRedis(key string)(string,error){
	cmd:=redisDb.Do("GET",key)
	result,err:=cmd.String()
	if err!=nil && err.Error()=="redis: nil" {
		return "",nil
	}else if err!=nil {
		log.Printf("GetValueFromRedis error,err:%+v\n",err)
		return "",err
	}
	return result,nil
}