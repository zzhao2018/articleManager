package logic

import (
	"articleManager/dao"
	"articleManager/model"
	"fmt"
)

//查询类型数据
func SearchAllType()([]*model.TypeInfo){
	return dao.SearchType()
}

//使用id查找类型数据
func SearchTypeById(typeI int)(*model.TypeInfo){
	return dao.SearchTypeById(typeI)
}

//更改时间
func AlterType(typeI int,hours int,mins int)error{
	timeS:=fmt.Sprintf("%d:%d",hours,mins)
	return dao.AlterTypeTime(typeI,timeS)
}

//删除类型数据
func DeleteType(typeI []int)error{
	//遍历并删除数据
	for _,typeEle:=range typeI {
		err:=dao.DeletaTypeById(typeEle)
		if err!=nil {
			return err
		}
	}
	return nil
}

//新增类型
func AddType(typeInfoList map[string][]int)error{
	//查询是否有重名
	for eleKey,typeEleArr:=range typeInfoList{
		var typeInfo=model.TypeInfo{}
		typeInfo.TypeName=eleKey
		typeInfo.SendTime=fmt.Sprintf("%d:%d",typeEleArr[0],typeEleArr[1])
		//检查是否同名
		ele:=dao.SearchTypeByName(typeInfo.TypeName)
		if ele!=nil&&ele.Id>0 {
			err:=fmt.Errorf("AddType rename type")
			return err
		}
		//插入数据
		dao.InsertType(&typeInfo)
	}
	return nil
}