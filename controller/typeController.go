package controller

import (
	"articleManager/logic"
	"articleManager/wxutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/*
  类别处理的controller
*/

//删除类别
func DeleteType(contextG *gin.Context){
	//获得删除的id
	var deletearr []int
	deleteL:=contextG.Request.FormValue(wxutil.C_ParamDelete)
	//解析
	err:=json.Unmarshal([]byte(deleteL),&deletearr)
	if err!=nil {
		log.Printf("DeleteType json Unmarshal error,err:%+v\n",err)
		wxutil.ResponseData(contextG,nil,err,0)
		return
	}
	//重置数据库
	err=logic.DeleteType(deletearr)
	if err!=nil {
		log.Printf("DeleteType logic.DeleteType error,err:%+v\n",err)
		wxutil.ResponseData(contextG,nil,err,0)
		return
	}
	//重置定时任务
	typeMap,err:=resertAfterAlter()
	if err!=nil {
		log.Printf("DeleteType resertAfterAlter error,err:%+v\n",err)
		wxutil.ResponseData(contextG,nil,err,0)
		return
	}
	wxutil.ResponseData(contextG,typeMap,nil,http.StatusOK)
}

//新增类别
func AddType(contextG *gin.Context){
	//获得删除的id
	var insertType map[string][]int
	insertL:=contextG.Request.FormValue(wxutil.C_ParamInsert)
	//解析
	err:=json.Unmarshal([]byte(insertL),&insertType)
	if err!=nil {
		log.Printf("AddType json Unmarshal error,err:%+v\n",err)
		wxutil.ResponseData(contextG,nil,err,0)
		return
	}
	//插入数据库
	err=logic.AddType(insertType)
	if err!=nil {
		log.Printf("AddType logic.DeleteType error,err:%+v\n",err)
		wxutil.ResponseData(contextG,nil,err,0)
		return
	}
	//重置定时任务
	typeMap,err:=resertAfterAlter()
	if err!=nil {
		log.Printf("AddType resertAfterAlter error,err:%+v\n",err)
		wxutil.ResponseData(contextG,nil,err,0)
		return
	}
	wxutil.ResponseData(contextG,typeMap,nil,http.StatusOK)
}

//删除、增加时重置定时任务
func resertAfterAlter()(map[string][]int,error){
	//获得剩余类型数据
	typeInfo:=logic.SearchAllType()
	typeMap:=make(map[string][]int,len(typeInfo))
	for _,typeEle:=range typeInfo  {
		mapKey:=strconv.Itoa(int(typeEle.Id))
		arrValue:=make([]int,0)
		//获得时和分
		timeS:=typeEle.SendTime
		timeArr:=strings.Split(timeS,":")
		for _,timeEleS:=range timeArr{
			timeEle,err:=strconv.ParseInt(timeEleS,10,64)
			if err!=nil {
				log.Println("resertAfterAlter conver error,err:%+v\n",err)
				return nil,err
			}
			arrValue=append(arrValue,int(timeEle))
		}
		//设置map
		typeMap[mapKey]=arrValue
	}
	//重启线程
	//抽取文案类型、时间
	useType,hours,mins,err:=getUserHoursMinArr(typeMap)
	if err!=nil {
		log.Printf("resertAfterAlter getUserHoursMinArr error,err:%+v\n",err)
		return nil,err
	}
	//重置发送时间
	err=resetTime(useType,hours,mins)
	if err!=nil {
		log.Printf("resertAfterAlter resetTime error,err:%+v\n",err)
		return nil,err
	}
	return typeMap,nil
}

//重置定时设置
func ReSetSendParam(contextG *gin.Context){
	//解析配置文件
	paramStr:=contextG.Request.FormValue(wxutil.C_ParamSeting)
	//解析前端json数据
	var paramArr map[string][]int
	err:=json.Unmarshal([]byte(paramStr),&paramArr)
	if err!=nil {
		log.Printf("ReSetSendParam Unmarshal error,err:%+v\n",err)
		wxutil.ResponseData(contextG,nil,err,wxutil.GetErrorCode(err.Error()))
		return
	}
	log.Printf("get paramarr :%+v\n",paramArr)
	//抽取文案类型、时间
	useType,hours,mins,err:=getUserHoursMinArr(paramArr)
	if err!=nil {
		log.Printf("ReSetSendParam getUserHoursMinArr error,err:%+v\n",err)
		wxutil.ResponseData(contextG,nil,err,wxutil.GetErrorCode(err.Error()))
		return
	}
	//更新数据库
	changeType:=getChangeType(useType,hours,mins)
	log.Printf("change type:%+v\n",changeType)
	for _,changeTypeEle:=range changeType  {
		err:=logic.AlterType(changeTypeEle[0],changeTypeEle[1],changeTypeEle[2])
		if err!=nil {
			log.Printf("resetTime alter type error,err:%+v\n",err)
		}
	}
	//重置发送时间
	err=resetTime(useType,hours,mins)
	if err!=nil {
		log.Printf("ReSetSendParam resetTime error,err:%+v\n",err)
		wxutil.ResponseData(contextG,nil,err,wxutil.GetErrorCode(err.Error()))
		return
	}
	//正常返回
	wxutil.ResponseData(contextG,paramArr,nil,http.StatusOK)
}


//获得类型列表
func GetTypeList(ctx *gin.Context){
	typeList:=logic.SearchAllType()
	wxutil.ResponseData(ctx,typeList,nil,http.StatusOK)
}
