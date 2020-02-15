package logic

import (
	"articleManager/dao"
	"log"
	"sync"
	"time"
)

//用以判断是否登录
var (
	loginMap map[string]bool=make(map[string]bool)
	mapLock  sync.Mutex
)

func init()  {
	go clearPairInTime()
}

//判断登录是否成功
func LoginAccess(password string) bool {
	userInfo := dao.GetPassWord()
	if password == userInfo.Password {
		return true
	}
	return false
}

//判断是否登录过
func UserIsLogin(key string, timeLong int) bool {
	mapLock.Lock()
	defer mapLock.Unlock()
	//字典中不存在，返回未登录
	if loginStatus, ok := loginMap[key]; ok == false {
		return false
		//字典中存在，但是键过期，返回false
	} else if loginStatus == false {
		return false
		//字典中存在，内存中显示未过期，查redis
	} else {
		value, err := dao.GetValueFromRedis(key)
		if err != nil {
			log.Printf("UserIsLogin GetValueFromRedis error,err:%+v\n", err)
			loginMap[key] = false
			return false
		}
		if value == "1" {
			//续期
			dao.SetValueToRedis(key, "1", timeLong)
			return true
		}
		loginMap[key] = false
		return false
	}
}

//设置redis时间
func SetUserTime(key string, value string, timeLen int) {
	mapLock.Lock()
	defer mapLock.Unlock()
	dao.SetValueToRedis(key, value, timeLen)
	loginMap[key]=true
}


//定期清理无用键
func clearPairInTime(){
	for {
		time.Sleep(time.Hour*24)
		go clearPair()
	}
}

//清理无用键值对
func clearPair(){
	mapLock.Lock()
	defer mapLock.Unlock()
	for k,v:=range loginMap  {
		if v==false {
			delete(loginMap,k)
		}
	}
}