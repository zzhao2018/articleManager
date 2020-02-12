package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"testing"
)


//测试yaml
func TestConf_GetDataBaseUrl(t *testing.T) {
	fileData,err:=ioutil.ReadFile("conf.yaml")
	var conf Conf
	if err!=nil {
		log.Printf("config file can not found,err:%+v\n",err)
		return
	}
	err=yaml.Unmarshal(fileData,&conf)
	if err!=nil {
		log.Printf("config unmarshal error,err:%+v\n",err)
		return
	}
	fmt.Printf("%s\n",conf.GetDataBaseUrl())
	fmt.Printf("data:%+v\n",conf.PhotoSavePath)
	fmt.Printf("data:%+v\n",conf.EmailAddr)
	fmt.Printf("data:%+v\n",conf.ToEmailAddr)
	fmt.Printf("data:%+v\n",conf.Subject)
}
