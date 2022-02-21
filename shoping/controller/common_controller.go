package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shoping/model"
	"shoping/param"
	"shoping/service"
	"shoping/tool"
	"strconv"
	"time"
	"unicode/utf8"
)

type CommonController struct {

}
var(
	cs=service.CommonService{}
	gs=service.GoodService{}
	ts=service.Tokenservice{}

)
func(c *CommonController)Route(engine *gin.Engine){
	engine.GET("/api/good/comments", c.getComments)
	engine.POST("/api/good/comment", c.postComment)
}
func(c *CommonController)getComments(ctx *gin.Context){
	goodId:=ctx.Query("good_id")
	if goodId==""{
		tool.RespErrorWithDate(ctx,"商品id不能为空")
		return
	}
	id,err:=strconv.ParseInt(goodId,10,64)
	if err!=nil{
		fmt.Println("获取商品id失败:",err)
		tool.RespInternalError(ctx)
		return
	}
	flag,err:=gs.JudgeGoodId(id)
	if err!=nil{
		fmt.Println("判断id错误:",err)
		tool.RespInternalError(ctx)
		return
	}
	if flag==false{
		tool.RespErrorWithDate(ctx,"商品id为空")
		return
	}
	commonSlice,err:=cs.GetCommonSlice(id)
	if err!=nil{
		fmt.Println("获取评论失败")
		tool.RespInternalError(ctx)
		return
	}
	if commonSlice==nil{
		commonSlice = []model.Comment{}
	}
	var commonParam param.CommentParam
	var newCommonSlice []param.CommentParam
	for _,commonModel:= range commonSlice{
		user,_:=us.GetUserInfo(commonModel.UserId)


		commonParam.Time = time.Now().Format("2006-01-02 15:04:05")
		commonParam.Id = commonModel.Id
		commonParam.Value = commonModel.Value
		commonParam.User.Username = user.Username
		commonParam.Likes = 0
		commonParam.GoodId = commonModel.GoodId
		newCommonSlice=append(newCommonSlice,commonParam)

	}
	tool.RespSuccessfulWithDate(ctx,newCommonSlice)

}
func(c *CommonController)postComment(ctx *gin.Context){
	goodid:=ctx.PostForm("good_id")
	if goodid==""{
		fmt.Println("商品id不能为空")
		tool.RespInternalError(ctx)
		return

	}
	id,err:=strconv.ParseInt(goodid,10,64)
	if err!=nil{
		fmt.Println("解析商品id失败")
		tool.RespInternalError(ctx)
		return

	}
	flag,err:=gs.JudgeGoodId(id)
	if err!=nil{
		fmt.Println("判断id服务出错")
		tool.RespInternalError(ctx)
		return
	}
	if flag==false{
		tool.RespErrorWithDate(ctx,"输入的商品id失效")
		return
	}
	//解析token
	token:=ctx.PostForm("token")
	claim,err:=st.ParamToken(token)
	if err!=nil{
		fmt.Println("解析token服务实现")
		tool.RespInternalError(ctx)
		return
	}
	//验证token
	flag=tool.CheckTokenErr(ctx,claim,err)
	if flag==false{
		return

	}
	user:=model.User{}
	comment:=ctx.PostForm("comment")
	if comment==""{
		tool.RespErrorWithDate(ctx,"评论不能为空")
		return
	}
	if utf8.RuneCountInString(comment) > 1024{
		tool.RespErrorWithDate(ctx,"评论内容过长")
	}
	var commentModel model.Comment
	commentModel.Value=comment
	commentModel.UserId=user.Id
	commentModel.GoodId=id
	//将评论插入数据库，并获得评论的id
	commentId,err:=cs.PostCommentSlice(commentModel)
	if err!=nil{
		fmt.Println("评论插入数据库失败")
		tool.RespInternalError(ctx)
		return
	}
	commentModel.Time=time.Now()
	commentModel.Likes=0
	commentModel.Id=commentId
	tool.RespSuccessfulWithDate(ctx,commentModel)

}
