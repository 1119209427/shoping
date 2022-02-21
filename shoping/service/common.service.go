package service

import (
	"shoping/dao"
	"shoping/model"
	"shoping/tool"
)

type CommonService struct{

}
var(
	cd=dao.CommonDao{tool.GetDB()}
)
func (cs *CommonService)GetCommonSlice(id int64)([]model.Comment,error){
	commonSlice,err:=cd.QueryById(id)
	if err!=nil{
		return nil, err
	}
	return commonSlice,nil
}
func(cs *CommonService)PostCommentSlice(comment model.Comment)(int64,error){
	id,err:=cd.InsertComment(comment)
	if err!=nil{
		return 0,err
	}
	return id,nil

}




