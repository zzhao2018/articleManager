package dao

import "articleManager/model"

//查询所有数据
func SearchType()([]*model.TypeInfo){
	var typeInfoL []*model.TypeInfo
	db.Where("delete_code=?",0).Find(&typeInfoL)
	return typeInfoL
}

//查出部分数据
func SearchTypeById(typeI int)*model.TypeInfo{
	var typeInfo model.TypeInfo
	db.Where("id=? and delete_code=?",typeI,0).First(&typeInfo)
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

//删除类别,将delete_code设置为1
func DeletaTypeById(typeI int)error{
	var typeModel=model.TypeInfo{
		Id:       int32(typeI),
	}
	db.Model(&typeModel).Update("delete_code",1)
	return nil
}

