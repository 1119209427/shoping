package service

import (
	"shoping/dao"
	"shoping/tool"
)

type OrderService struct {

}
var (
	od =dao.OrderDao{tool.GetDB()}
)
func(os *OrderService)GetGoodid(uid int64,state string)([]int64,error){
	gIdSlice,err:=od.GetGoodUid(uid,state)
	if err!=nil{
		return nil,err
	}
	return gIdSlice,nil

}
func(os *OrderService)ChangeState(uid int64,newType string)error{

	err:=od.ChangeState(uid,newType)
	if err!=nil{
		return err
	}
	return nil


}
