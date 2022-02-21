package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shoping/tool"
)

type TokenController struct {

}
func(tc *TokenController)Route(engine *gin.Engine){
	engine.GET("/api/verify/token",tc.getToken)
}
func(tc *TokenController)getToken(ctx *gin.Context){
	refreshToken:=ctx.Query("refreshToken")
	//判断refreshtoken的状态
	claim,err:=ts.ParamRefreshToken(refreshToken)
	if err!=nil{
		if err.Error()[:16]=="token is expired"{
			tool.RespErrorWithDate(ctx,"refreshToken失效")
			return
		}
		fmt.Println("getTokenParseTokenErr:", err)
		tool.RespInternalError(ctx)
		return
	}
	if claim.Type=="ERR"{
		tool.RespErrorWithDate(ctx,"refreshToken错误或服务出错")
		return
	}
	//根据id更新用户信息
	user,err:=us.GetUserInfo(claim.User.Id)
	if err!=nil{
		fmt.Println("更新用户信息失败")
		tool.RespInternalError(ctx)
		return
	}
	//创建新的token
	newToken,err:=ts.CreateToken(user,120,"TOKEN")
	if err!=nil{
		fmt.Println("创建token失败")
		tool.RespInternalError(ctx)
		return

	}
	tool.RespSuccessfulWithDate(ctx,newToken)

}
