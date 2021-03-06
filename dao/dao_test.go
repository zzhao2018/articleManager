package dao

import (
	"articleManager/model"
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
	InitRedis(config)
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

func TestGetPassWord(t *testing.T) {
	t.Logf("%+v",GetPassWord())
}

func TestSetValueToRedis(t *testing.T) {
	SetValueToRedis("test","value",60)
}

func TestGetValueFromRedis(t *testing.T) {
	data,err:=GetValueFromRedis("test")
	t.Logf("data:%+v\n",data)
	t.Logf("err:%+v\n",err)
}

func TestDeletaTypeById(t *testing.T) {
	err:=DeletaTypeById(1)
	if err!=nil {
		t.Logf("error:%+v\n",err)
		return
	}
}


func TestInsertType(t *testing.T) {
	InsertType(
		&model.TypeInfo{
			Id:          0,
			TypeName:    "测试type",
			SendTime:    "12:30",
			Delete_code: 0,
		})
}

func TestSearchTypeByName(t *testing.T) {
	data:=SearchTypeByName("关系")
	t.Logf("data:%+v\n",data)
}