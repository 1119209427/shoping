package service

import (
	"github.com/gin-gonic/gin"
	"shoping/dao"
	"shoping/model"
	"shoping/tool"
)

type CartService struct {

}
var(
	d=dao.CartDao{tool.GetDb()}
)
func(c *CartService)JudgeCartId(cartId int64)(bool,error){
	_,err:=d.QueryByCid(cartId)
	if err!=nil{
		if err.Error()=="ql: no rows in result set"{
			return false, nil
		}
		return false, err
	}
	return true,nil


}
func(c CartService)GetGoodInfoInRedis(ctx *gin.Context,key string)([]model.Good,error){
	var goodSlice []model.Good
	good,err:=rd.RedisGetValueShop(ctx,key)
	if err!=nil{
		return nil, err
	}
	goodSlice=append(goodSlice,good)
	return goodSlice,nil

}
