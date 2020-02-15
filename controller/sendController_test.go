package controller

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
	"time"
	"articleManager/conf"
)

func TestSetTickWork(t *testing.T) {
	ctx,cancle:=context.WithCancel(context.TODO())
	useType:=[]int{1,2,3}
	hours:=[]int{0,0,0}
	mins:=[]int{40,41,41}
	go SetTickWork(ctx,useType,hours,mins)
	time.Sleep(time.Minute*4)
	cancle()
	for  {
		time.Sleep(time.Second*10)
		fmt.Printf("wati........\n")
	}
}



func TestDoJob(t *testing.T) {

}

func initYaml()*conf.Conf{
	//读取文件信息
	fileData,err:=ioutil.ReadFile(`../`+conf.C_ConfFilePath)
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

func init(){
	conf:=initYaml()
	SetConfig(conf)
}