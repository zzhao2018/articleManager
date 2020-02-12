package dao

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"testing"
	"articleManager/conf"
)

//解析配置文件
func initYaml()*conf.Conf{
	//读取文件信息
	fileData,err:=ioutil.ReadFile(`../conf/conf.yaml`)
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

var (
	config *conf.Conf
)

func init(){
	config=initYaml()
	InitDataBase(config)
}

func TestSearchArticle(t *testing.T) {
	result:=SearchArticle()
	for _,ele:=range result {
		log.Printf("data:%+v\n",ele)
	}
}

func TestSearchArticleById(t *testing.T) {
	result:=SearchArticleById(4)
	t.Logf("data:%+v\n",result)
}

func TestSearchArticleByRandom(t *testing.T) {
	result:=SearchArticleByRandom(1)
	t.Logf("data:%+v\n",result)
}

func TestSearchType(t *testing.T) {
	result:=SearchType()
	for _,ele:=range result  {
		t.Logf("data:%+v\n",ele)
	}
}

func TestSearchTypeById(t *testing.T) {
	result:=SearchTypeById(2)
	t.Logf("data:%+v\n",result)
}

func TestAlterTypeTime(t *testing.T) {
	AlterTypeTime(2,"0:30")
}