package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"time"
	"articleManager/logic"
	"articleManager/wxutil"
)

//增加文章
func AddArticle(context *gin.Context){
	//判断是否登录
	if checkLoginStatus(context)==false {
		wxutil.ResponseData(context,"",fmt.Errorf("unlogin"),0)
		return
	}
	//从context中获取数据
	paramMap:=make(map[string]interface{})
	//得到数据
	form,err:=context.MultipartForm()
	if err!=nil {
		log.Printf("AddArticle MultipartForm error,err:%+v\n",err)
		errorCode:= wxutil.GetErrorCode(err.Error())
		wxutil.ResponseData(context,nil,err,errorCode)
		return
	}
	//获取参数
	for k,v:=range form.Value  {
		paramMap[k]=v[0]
	}
	//获取文件
	fileList:=make([]string,0)
	photos:=form.File[wxutil.C_PhotoName]
	for _,v:=range photos{
		//设置保存路径
		filename:=v.Filename
		nowT:=time.Now()
		dst:=fmt.Sprintf("%s%s%s",config.PhotoSavePath,string(filepath.Separator),
			fmt.Sprintf("%d_%s", nowT.Nanosecond(),filename))
		//上传图片
		err:=context.SaveUploadedFile(v,dst)
		if err!=nil {
			log.Printf("AddArticle SaveUploadedFile error,err:%+v\n",err)
			wxutil.ResponseData(context,nil,err, wxutil.GetErrorCode(err.Error()))
			return
		}
		fileList= append(fileList, dst)
	}
	paramMap[wxutil.C_PhotoName]=fileList
	//保存数据库
	err=logic.AddArticle(paramMap)
	if err!=nil {
		log.Printf("AddArticle logic AddArticle error,err:%+v\n",err)
		wxutil.ResponseData(context,paramMap,err, wxutil.GetErrorCode(err.Error()))
		return
	}
	wxutil.ResponseData(context,paramMap,nil,http.StatusOK)
}

//获得类型列表
func GetTypeList(ctx *gin.Context){
	if checkLoginStatus(ctx)==false {
		wxutil.ResponseData(ctx,"",fmt.Errorf("unlogin"),0)
		return
	}
	typeList:=logic.SearchAllType()
	wxutil.ResponseData(ctx,typeList,nil,http.StatusOK)
}