package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"shoping/model"
	"shoping/service"
	"shoping/tool"
	"strconv"
)

type CartController struct {

}
func(c *CartController)Route(engine *gin.Engine){
	engine.GET("/api/cart/settlement",c.goodSettlement)//购物车中的商品结算
	engine.DELETE("/api/cart/delete",c.goodDelete)//购物车中的商品删除
}
var(
	s=service.CartService{}
)
func(c *CartController)goodSettlement(ctx *gin.Context){
	//获取购物车的id
	cartId:=ctx.Query("cartid")
	if cartId==""{
		tool.RespErrorWithDate(ctx,"购物车id不能为空")
		return
	}
	cId,err:=strconv.ParseInt(cartId,10,64)
	if err!=nil{
		fmt.Println("解析失败")
		tool.RespInternalError(ctx)
		return
	}
	flag,err:=s.JudgeCartId(cId)
	if err!=nil{
		fmt.Println("判断失败")
		tool.RespInternalError(ctx)
		return

	}
	if flag==false{
		tool.RespErrorWithDate(ctx,"购物车id出错")
		return
	}

	//验证token
	token:=ctx.Query("token")
	claim,err:=ts.ParamToken(token)
	if err!=nil{
		fmt.Println("解析token失败")
		tool.RespInternalError(ctx)
		return
	}
	flag=tool.CheckTokenErr(ctx,claim,err)
	if flag==false{
		return
	}
	//更新用户信息
	//userInfo:=claim.User
	//var cart model.ShoppingCart
	//根据购物车id获取redis中的商品信息
	goodSlice,err:=s.GetGoodInfoInRedis(ctx,string(cId))
	if err!=nil{
		fmt.Println("获取购物车信息失败")
		tool.RespInternalError(ctx)
		return
	}
	for _,v :=range goodSlice{
		fmt.Println(v)


	}




}
func(c *CartController)goodDelete(ctx *gin.Context){

}

