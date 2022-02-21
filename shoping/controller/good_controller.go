package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shoping/tool"
	"strconv"
)

type GoodController struct {

}

func(gc *GoodController)Route(engine *gin.Engine){
	engine.POST("/api/good/follow",gc.followGood)
	engine.POST("/api/good/shopping cart",gc.addtoShoppingCart)
	engine.POST("/api/good/likes",gc.increaseLikes)
	engine.GET("/api/good/info/:goodid",gc.getGoodinfo)
	engine.GET("/api/good/likes",gc.getGoodLike)

}
//加入购物车
func(gc *GoodController)addtoShoppingCart(ctx *gin.Context){
	//获取商品id
	goodId:=ctx.PostForm("goodid")
	if goodId==""{
		tool.RespErrorWithDate(ctx,"商品id不能为空")
		return
	}
	gid,err:=strconv.ParseInt(goodId,10,64)
	if err!=nil{
		fmt.Println("解析参数失败")
		tool.RespInternalError(ctx)
		return
	}
	//验证id
	flag,err:=gs.JudgeGoodId(gid)
	if err!=nil{
		fmt.Println("验证商品id错误")
		tool.RespInternalError(ctx)
		return
	}
	if flag==false{
		tool.RespErrorWithDate(ctx,"商品id错误")
		return
	}
	//解析token
	token:=ctx.PostForm("token")
	claim,err:=st.ParamToken(token)
	if err!=nil{
		fmt.Println("解析token失败")
		tool.RespInternalError(ctx)
		return
	}
	//验证token
	flag=tool.CheckTokenErr(ctx,claim,err)
	if flag==false{
		return

	}
	//更新用户信息
	user:=claim.User
	//将商品信息存入redis
	good,err:=gs.GetGoodInfo(gid)
	err=gs.GoodInShoppingCart(ctx,user.CartId,good)
	if err!=nil{
		fmt.Println("商品存入购物车出错")
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx,"商品成功存入购物车")

}
//获取喜爱状态
func(gc *GoodController)getGoodLike(ctx *gin.Context){
	//获取商品的id
	Goodid:=ctx.Query("goodid")
	if Goodid==""{
		tool.RespErrorWithDate(ctx,"商品id不能为空")
		return
	}
	gid,err:=strconv.ParseInt(Goodid,10,64)
	if err!=nil{
		fmt.Println("解析参数失败")
		tool.RespInternalError(ctx)
		return
	}
	//验证id
	flag,err:=gs.JudgeGoodId(gid)
	if err!=nil{
		fmt.Println("验证商品id错误")
		tool.RespInternalError(ctx)
		return
	}
	if flag==false{
		tool.RespErrorWithDate(ctx,"商品id错误")
		return
	}
	//解析token
	token:=ctx.Query("token")
	claim,err:=st.ParamToken(token)
	if err!=nil{
		fmt.Println("解析token失败")
		tool.RespInternalError(ctx)
		return
	}
	//验证token
	flag=tool.CheckTokenErr(ctx,claim,err)
	if flag==false{
		return

	}
	//更新用户信息
	user:=claim.User
	//获取喜爱状态
	flag,err=gs.GetLikesStatus(gid,user.Id)
	if err!=nil{
		fmt.Println("获取喜爱状态失败")
		tool.RespInternalError(ctx)
		return
	}
	if flag==false{
		tool.RespSuccessfulWithDate(ctx,"已经喜爱")
	}
	tool.RespSuccessfulWithDate(ctx,"添加喜欢成功")

}
//增加喜爱数
func(gc *GoodController)increaseLikes(ctx *gin.Context){
	//获取商品的id
	goodId:=ctx.PostForm("goodid")
	if goodId==""{
		fmt.Println("获取商品id失败")
		tool.RespInternalError(ctx)
		return
	}
	//转换用户id
	Gid,err:=strconv.ParseInt(goodId,10,64)
	if err!=nil{
		fmt.Println("解析失败")
		tool.RespInternalError(ctx)
		return
	}
	//判断id是否真确
	flag,err:=gs.JudgeGoodId(Gid)
	if err!=nil{
		fmt.Println("判断id失败")
		tool.RespInternalError(ctx)
		return
	}
	if flag==false{
		tool.RespErrorWithDate(ctx,"商品id错误")
		return
	}
	//解析token
	token:=ctx.PostForm("token")
	claim,err:=ts.ParamToken(token)
	flag=tool.CheckTokenErr(ctx,claim,err)
	if flag==false{
		return
	}
	//根据id增加点赞数
	//1获取点赞状态
	flag,err=gs.GetLikesStatus(Gid,claim.User.Id)
	if err!=nil{
		fmt.Println("获取点赞状态失败")
		tool.RespInternalError(ctx)
		return
	}
	//增加点赞
	err=gs.IncreaseLikes(flag,Gid,claim.User.Id)
	if err!=nil{
		fmt.Println("增加点赞失败")
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx,!flag)
}
//关注商品
func(gc *GoodController)followGood(ctx *gin.Context){
	//通过验证token检查用户是否登录
	//获取token
	token:=ctx.PostForm("token")
	if token==""{
		tool.RespErrorWithDate(ctx,"token不能为空")
		return

	}
	//解析token
	claim,err:=st.ParamToken(token)
	if err!=nil{
		fmt.Println("解析token失败",err)
		tool.RespInternalError(ctx)
		return

	}
	//验证token
	flag:=tool.CheckTokenErr(ctx,claim,err)
	if flag==false{
		return
	}
	user:=claim.User
	followId:=user.Id//获取要关注商品用户的id
	goodid:=ctx.PostForm("goodid")//获取被关注商品的id
	if goodid==""{
		fmt.Println("获取商品id失败")
		tool.RespInternalError(ctx)
		return
	}
	id,err:=strconv.ParseInt(goodid,10,64)
	if err!=nil{
		fmt.Println("解析失败")
		tool.RespInternalError(ctx)
		return
	}

	//检查id是否合法
	flag,err=gs.JudgeGoodId(id)
	if err!=nil{
		fmt.Println("判断商品id失败")
		tool.RespInternalError(ctx)
		return
	}
	if flag==false{
		tool.RespErrorWithDate(ctx,"商品id非法")
		return
	}
	//通过id获取商品详细

	flag,err=gs.GetFollowStatus(id,followId)
	if err!=nil{
		fmt.Println("获取关注状态失败")
		tool.RespInternalError(ctx)
		return
	}
	//更新关注状态
	err=gs.SolveFollow(flag,followId,id)
	if err!=nil{
		fmt.Println("更新关注状态失败")
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx,!flag)

}
func(gc *GoodController)getGoodinfo(ctx *gin.Context){
	goodid:=ctx.Param("goodid")
	if goodid==""{
		fmt.Println("获取商品id服务失效")
		tool.RespInternalError(ctx)
		return
	}
	id,err:=strconv.ParseInt(goodid,10,64)
	if err!=nil{
		fmt.Println("参数解析失败")
		tool.RespInternalError(ctx)
		return
	}
	//判断goodid
	flag,err:=gs.JudgeGoodId(id)
	if err!=nil{
		fmt.Println("判断商品id服务失效")
		tool.RespInternalError(ctx)
		return
	}
	if flag==false{
		tool.RespErrorWithDate(ctx,"商品id无效")
		return
	}
	//通过id获得商品信息
	good,err:=gs.GetGoodInfo(id)
	if err!=nil{
		fmt.Println("获取商品信息服务失效")
		tool.RespInternalError(ctx)
		return
	}
	result:=tool.ObjToMap(good)
	tool.RespSuccessfulWithDate(ctx,result)




}
