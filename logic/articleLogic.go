package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
	"articleManager/dao"
	"articleManager/model"
	"articleManager/wxutil"
)

//添加文章到数据库
func AddArticle(paramMap map[string]interface{}) error{
	//获得文章信息
	context,ok:=paramMap[wxutil.C_ContextName]
	if ok==false {
		log.Printf("AddArticle no context error,data:%+v\n",paramMap)
		err:=fmt.Errorf("AddArticle no context error")
		return err
	}
	contextS,ok:=context.(string)
	if ok==false {
		log.Printf("AddArticle change context error,data:%+v\n",context)
		err:=fmt.Errorf("AddArticle change context error")
		return err
	}
	//获得图片信息
	var photoS string
	photo:=paramMap[wxutil.C_PhotoName]
	if photo!=nil && photo!="" {
		photoL,ok:=photo.([]string)
		if ok==false {
			log.Printf("AddArticle change photo error,data:%+v\n",context)
			err:=fmt.Errorf("AddArticle change photo error")
			return err
		}
		photoArr,err:=json.Marshal(photoL)
		if err!=nil {
			log.Printf("AddArticle Marshal photo error,data:%+v,err:%+v\n",photoL,err)
			return err
		}
		photoS=string(photoArr)
	}

	//获得类型信息
	typeInfo,ok:=paramMap[wxutil.C_TypeName]
	if ok==false {
		log.Printf("AddArticle no type error,data:%+v\n",paramMap)
		err:=fmt.Errorf("AddArticle no type error")
		return err
	}
	typeS,ok:=typeInfo.(string)
	if ok==false {
		log.Printf("AddArticle type change error,data:%+v\n",typeInfo)
		err:=fmt.Errorf("AddArticle type change error")
		return err
	}
	typeI,err:=strconv.ParseInt(typeS,10,32)
	if err!=nil {
		log.Printf("AddArticle type change to int error,data:%+v\n",typeS)
		return err
	}
	//插入数据
	articleInfo:=&model.ArticleInfo{
		Article_context: contextS,
		Type:            int32(typeI),
		Photo:           photoS,
		Insert_time:time.Now(),
		Update_time:time.Now(),
	}
	dao.InsertArticle(articleInfo)
	return nil
}

//随机获取一条消息
func SearchRandomArticle(typeI int)(*model.ArticleInfo){
	return dao.SearchArticleByRandom(typeI)
}