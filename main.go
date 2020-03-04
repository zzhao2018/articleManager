package main

import (
	"articleManager/conf"
	"articleManager/controller"
	"articleManager/dao"
	"articleManager/wxutil"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//解析配置文件
func initYaml()*conf.Conf{
	//读取文件信息
	fileData,err:=ioutil.ReadFile(conf.C_ConfFilePath)
	var conf conf.Conf
	if err!=nil {
		panic("config file can not found")
	}
	//解码信息
	err=yaml.Unmarshal(fileData,&conf)
	if err!=nil {
		panic("config unmarshal error")
	}
	return &conf
}

//初始化参数
func initProject()(*conf.Conf,error){
	//初始化配置文件
	conf:=initYaml()
	controller.SetConfig(conf)
	//初始化数据库
	err:=dao.InitDataBase(conf)
	if err!=nil {
		log.Printf("initProject InitDataBase error,err:%+v\n",err)
		return nil,err
	}
	//初始化redis
	dao.InitRedis(conf)
	//初始化id生成器
	err=wxutil.InitSonyFlake(0)
	if err!=nil {
		log.Printf("initProject InitSonyFlake error,err:%+v\n",err)
		return nil,err
	}
	return conf,nil
}

func main() {
	//初始化参数
	conf,err:=initProject()
	if err!=nil {
		log.Printf("initProject error,err:%+v\n",err)
		return
	}
	//开启服务
	engi:=gin.Default()
	//使用https
	engi.Use(controller.HttpSHandler())
	engi.POST("/addArticle",controller.LoginMiddleWare,controller.AddArticle)
	engi.POST("/alterParam",controller.LoginMiddleWare,controller.ReSetSendParam)
	engi.GET("/getTypeList",controller.LoginMiddleWare,controller.GetTypeList)
	engi.POST("/deleteType",controller.LoginMiddleWare,controller.DeleteType)
	engi.POST("/login",controller.LoginIn)
	engi.RunTLS(":8089",conf.PemPath,conf.SslPath)
}

