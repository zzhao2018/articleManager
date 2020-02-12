package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
	"articleManager/logic"
	"articleManager/wxutil"
)

//设置单个定时任务
func tickWork(ctx context.Context, useType int, hours int, mins int) {
	//设置时间
	nowT := time.Now()
	timeStamp := time.Date(nowT.Year(), nowT.Month(), nowT.Day(), hours, mins, 0, 0, nowT.Location())
	timeTick := time.NewTimer(timeStamp.Sub(nowT))
	//开始任务
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("close thread.....")
			return
		case <-timeTick.C:
			//更新时间
			nowT = time.Now()
			timeStamp = time.Date(nowT.Year(), nowT.Month(), nowT.Day()+1, hours, mins, 0, 0, nowT.Location())
			timeTick = time.NewTimer(timeStamp.Sub(nowT))
			//运行程序
			job := NewTickJob(useType)
			go job.Run()
		}
	}
}

//设置所有定时任务
func SetTickWork(ctx context.Context, useType []int, hours []int, mins []int) {
	if len(useType) <= 0 || len(hours) <= 0 || len(mins) <= 0 {
		return
	}
	for i := 0; i < len(useType); i++ {
		go tickWork(ctx, useType[i], hours[i], mins[i])
	}
}


//重置定时设置
func ReSetSendParam(contextG *gin.Context){
	//解析配置文件
	paramStr:=contextG.Query(wxutil.C_ParamSeting)
	var paramArr map[string][]int
	err:=json.Unmarshal([]byte(paramStr),&paramArr)
	if err!=nil {
		log.Printf("ReSetSendParam Unmarshal error,err:%+v\n",err)
		wxutil.ResponseData(contextG,nil,err,wxutil.GetErrorCode(err.Error()))
		return
	}
	log.Printf("get paramarr :%+v\n",paramArr)
	useType:=make([]int,0)
	hours:=make([]int,0)
	mins:=make([]int,0)
	for k,v:=range paramArr{
		//解析类别
		kEle,err:=strconv.ParseInt(k,10,32)
		if err!=nil {
			log.Printf("ReSetSendParam ParseInt error,err:%+v\n",err)
			wxutil.ResponseData(contextG,nil,err,wxutil.GetErrorCode(err.Error()))
			return
		}
		useType= append(useType, int(kEle))
		hours=append(hours,v[0])
		mins=append(mins,v[1])
	}
	log.Printf("useT:%+v hours:%+v mins:%+v\n",useType,hours,mins)
	//加锁设置定时线程
	sendLock.Lock()
	defer sendLock.Unlock()
	//取消先前任务
	if defaultcCancleFunc !=nil{
		defaultcCancleFunc()
	}
	//更新数据库
	changeType:=getChangeType(useType,hours,mins)
	log.Printf("change type:%+v\n",changeType)
	for _,changeTypeEle:=range changeType  {
		err=logic.AlterType(changeTypeEle[0],changeTypeEle[1],changeTypeEle[2])
		if err!=nil {
			log.Printf("ReSetSendParam alter type error,err:%+v\n",err)
		}
	}
	//设置当前任务
	SetDefaultSendParam(useType,hours,mins)
	go SetTickWork(defauleCtx,defaultUserType,defaultHour,defaultMin)
	wxutil.ResponseData(contextG,paramArr,nil,http.StatusOK)
}

//新获得和原本数据的差集
func getChangeType(useType []int,hours []int,mins []int)[][]int{
	typeMap:=make(map[int][]int)
	resultL:=make([][]int,0)
	//抽取默认字典
	for i:=0;i<len(defaultUserType) ;i++  {
		typeMap[defaultUserType[i]]=[]int{hours[i],mins[i]}
	}
	for i:=0;i<len(useType) ;i++  {
		//判断是否在里头
		timeEle,ok:=typeMap[useType[i]]
		if ok==false {
			midList:=make([]int,3)
			midList[0]=useType[i]
			midList[1]=hours[i]
			midList[2]=mins[i]
			resultL=append(resultL,midList)
		}else{
			if timeEle[0]==hours[i] && timeEle[1]==mins[i] {
				continue
			}else {
				midList:=make([]int,3)
				midList[0]=useType[i]
				midList[1]=hours[i]
				midList[2]=mins[i]
				resultL=append(resultL,midList)
			}
		}
	}
	return resultL
}