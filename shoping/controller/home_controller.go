package controller
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shoping/model"
	"shoping/service"
	"shoping/tool"
)

type HomeController struct {

}
var(
	hs=service.HomeService{}
)
func(h *HomeController)Route(engine *gin.Engine){
	engine.GET("/api/home/search", h.search)
	engine.GET("/api/home/sections", h.getSections)

}
func(h *HomeController)getSections(ctx *gin.Context){
	typeSlice:=[]string{"01%","02%","03%","04%","05%","06%","07%","information","08%","09%","10%","11%","13%","15%"}
	var Date []model.Section
	for i,channelType:=range typeSlice{
		var section model.Section
		if channelType=="information"{
			fmt.Println(i)
			//京东快报

			randSlice, rankSlice, err :=hs.GetChannelInfo()
			if err!=nil{
				fmt.Println("获取京东快报失败")
				tool.RespInternalError(ctx)
				return
			}
			section.List=randSlice
			section.Rank=rankSlice
			if section.List==nil{
				section.List=[]model.Good{}
			}
			if section.List==nil{
				section.Rank=[]model.Good{}
			}
			Date=append(Date,section)
			continue
		}
		randSlice, rankSlice, err :=hs.GetChannelGood(channelType)
		if err!=nil{
			fmt.Println("获取商品路径失败")
			tool.RespInternalError(ctx)
			return
		}
		section.List=randSlice
		section.Rank=rankSlice
		if section.List==nil{
			section.List=[]model.Good{}
		}
		if section.List==nil{
			section.Rank=[]model.Good{}
		}
		Date=append(Date,section)

	}
	fmt.Println(Date)
	tool.RespSuccessfulWithDate(ctx,Date)




}
func(h *HomeController)search(ctx *gin.Context){
	keyWord:=ctx.Query("key_word")
	if keyWord==""{
		fmt.Println("关键词不能为空")
		tool.RespInternalError(ctx)
		return
	}
	merchantWithGood,err:=hs.Search(keyWord)
	if err!=nil{
		fmt.Println("查询服务失效:",err)
		tool.RespInternalError(ctx)
		return
	}
	if merchantWithGood==nil{
		merchantWithGood=[]model.MerchantWithGood{}
	}
	tool.RespSuccessfulWithDate(ctx,merchantWithGood)

}

