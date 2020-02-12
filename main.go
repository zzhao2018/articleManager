package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"articleManager/conf"
	"articleManager/controller"
	"articleManager/dao"
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
func initProject()error{
	//初始化配置文件
	conf:=initYaml()
	controller.SetConfig(conf)
	//初始化数据库
	err:=dao.InitDataBase(conf)
	if err!=nil {
		log.Printf("initProject InitDataBase error,err:%+v\n",err)
		return err
	}
	return nil
}

func main() {
	//初始化参数
	err:=initProject()
	if err!=nil {
		log.Printf("initProject error,err:%+v\n",err)
		return
	}
	//开启服务
	engi:=gin.Default()
	engi.POST("/addArticle",controller.AddArticle)
	engi.POST("/alterParam",controller.ReSetSendParam)
	engi.GET("/getTypeList",controller.GetTypeList)

	engi.Run(":8088")
}