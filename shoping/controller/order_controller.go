package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shoping/model"
	"shoping/param"
	"shoping/service"
	"shoping/tool"
	"strconv"
)

type OrderController struct {

}
func (oc *OrderController)Route(engine *gin.Engine){
	engine.GET("/api/order/good",oc.getGoodInfo)//获取订单中的商品信息
	engine.PUT("/api/order/change",oc.changeState)//改变订单状态
}
func(oc *OrderController)getGoodInfo(ctx *gin.Context){
	//获取订单状态
	state:=ctx.Query("state")
	//获取用户的信息
	uid:=ctx.Query("uid")
	if uid ==""{
		fmt.Println("用户id不能为空")
		tool.RespInternalError(ctx)
		return
	}
	id,err:=strconv.ParseInt(uid,10,64)
	if err!=nil{
		fmt.Println("解析失败")
		tool.RespInternalError(ctx)
		return
	}
	flag,err:=us.JudgeUid(id)
	if err!=nil{
		fmt.Println("判断用户id失败")
		tool.RespInternalError(ctx)
		return
	}
	if flag==false{
		tool.RespErrorWithDate(ctx,"用户id错误")
		return
	}
	//通过id,订单状态找到订单表中的商品id,通过商品id获取商品详细信息
	var goodInfoSlice []model.Good
	var os service.OrderService
	gIdSlice,err:=os.GetGoodid(id,state)
	if err!=nil{
		fmt.Println("获取商品id失败")
		tool.RespInternalError(ctx)
		return
	}
	for _,gid:=range gIdSlice{
		goodInfo,err:=gs.GetGoodInfo(gid)
		if err!=nil{
			fmt.Println("获取商品信息失败")
			tool.RespInternalError(ctx)
			return
		}
		goodInfoSlice=append(goodInfoSlice,goodInfo)
	}
	tool.RespSuccessfulWithDate(ctx,goodInfoSlice)




}
func(oc *OrderController)changeState(ctx *gin.Context){
	//解析表单
	var orderParma param.ChangeOrder
	err:=ctx.ShouldBind(&orderParma)
	if err!=nil{
		fmt.Println("解析失败")
		tool.RespInternalError(ctx)
		return
	}
	if orderParma.Token==""{
		tool.RespErrorWithDate(ctx, "NO_TOKEN_PROVIDED")
		return
	}
	//解析token
	claim,err:=ts.ParamToken(orderParma.Token)
	if err!=nil{
		fmt.Println("解析失败")
		tool.RespInternalError(ctx)
		return
	}
	flag:=tool.CheckTokenErr(ctx,claim,err)
	if flag==false{
		return
	}
	user:=claim.User
	//改变状态
	newType:=ctx.PostForm("new_type")
	var os service.OrderService
	err=os.ChangeState(user.Id,newType)
	if err!=nil{
		fmt.Println("改变订单状态失败")
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx,"修改订单状态成功")

}
