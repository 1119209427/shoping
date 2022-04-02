package service

import (
	"log"
	"shoping/dao"
	"shoping/model"
	"shoping/tool"
)

type CommonService struct {
}

var (
	cd = dao.CommonDao{tool.GetDb()}
)

func (cs *CommonService) GetCommonSlice(id int64) ([]model.Comment, error) {
	db := tool.GetGormDb()
	var comment []model.Comment
	var user model.User
	if err := db.Where("id", id).Preload("User", user).Find(&comment).Error; err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	/*commonSlice,err:=cd.QueryById(id)
	if err!=nil{
		return nil, err
	}*/
	return comment, nil
}
func (cs *CommonService) PostCommentSlice(comment model.Comment) (int64, error) {
	db := tool.GetGormDb()

	if err := db.Delete(&comment).Error; err != nil {
		return 0, err
	}
	/*id,err:=cd.InsertComment(comment)
	if err!=nil{
		return 0,err
	}*/
	return comment.Id, nil

}
