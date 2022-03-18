package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shoping/service"
	"shoping/tool"
)

type OauthController struct {

}

func(c *OauthController)Route(engine *gin.Engine){
	engine.GET("api/login-github",c.LoginGithub)
	engine.POST("https://github.com/login/oauth/access_token")
}
func(c *OauthController)LoginGithub(ctx *gin.Context){
	code:=ctx.Query("code")
	if code==""{
		tool.RespInternalError(ctx)
		fmt.Println("获取code失败")
		return
	}
	//code 通过 github.com OAuth API 换取 token
	// token 根据GitHub 开发API接口获取用户信息 githubUser
	gu,err:=service.FetchGithubUser(code)
	if err!=nil{
		tool.RespErrorWithDate(ctx,"token错误")
		return
	}
	tool.RespSuccessfulWithDate(ctx,gu)







}
