package logic

import (
	"fmt"
	"articleManager/dao"
	"articleManager/model"
)

//查询类型数据
func SearchAllType()([]*model.TypeInfo){
	return dao.SearchType()
}

func SearchTypeById(typeI int)(*model.TypeInfo){
	return dao.SearchTypeById(typeI)
}

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