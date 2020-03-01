package logic

import (
	"articleManager/dao"
	"articleManager/wxutil"
	"log"
	"time"
)

//用以判断是否登录
var (
	loginMap =wxutil.NewSafeMap()
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
	//字典中不存在，返回未登录
	if loginStatus, ok := loginMap.Get(key); ok == false {
		return false
		//字典中存在，但是键过期，返回false
	} else if loginStatus == false {
		return false
		//字典中存在，内存中显示未过期，查redis
	} else {
		value, err := dao.GetValueFromRedis(key)
		if err != nil {
			log.Printf("UserIsLogin GetValueFromRedis error,err:%+v\n", err)
			loginMap.Put(key,false)
			return false
		}
		if value == "1" {
			//续期
			dao.SetValueToRedis(key, "1", timeLong)
			return true
		}
		loginMap.Put(key,false)
		return false
	}
}

//设置redis时间
func SetUserTime(key string, value string, timeLen int) {
	loginMap.Put(key,true)
	dao.SetValueToRedis(key, value, timeLen)
}


//定期清理无用键
//此处可优化
//每日3点30开始清理
func clearPairInTime(){
	for {
		timeN:=time.Now()
		doTime:=time.Date(timeN.Year(),timeN.Month(),timeN.Day()+1,3,30,0,0,timeN.Location())
		time.Sleep(doTime.Sub(timeN))
		go func() {
			loginMap.ClearPair()
		}()
	}
}
