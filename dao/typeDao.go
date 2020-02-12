package dao

import "articleManager/model"

//查询所有数据
func SearchType()([]*model.TypeInfo){
	var typeInfoL []*model.TypeInfo
	db.Find(&typeInfoL)
	return typeInfoL
}

//查出部分数据
func SearchTypeById(typeI int)*model.TypeInfo{
	var typeInfo model.TypeInfo
	db.Where("id=?",typeI).First(&typeInfo)
	return &typeInfo
}

//更改时间
func AlterTypeTime(typeI int,timeS string)error{
	var typeModel=model.TypeInfo{
		Id:       int32(typeI),
	}
	db.Model(&typeModel).Update("send_time",timeS)
	return nil
}