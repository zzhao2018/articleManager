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
