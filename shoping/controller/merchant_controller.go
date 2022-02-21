package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shoping/model"
	"shoping/service"
	"shoping/tool"
	"strconv"
)

type MerchantController struct {

}
func(mc *MerchantController)Route(engine *gin.Engine){
	engine.GET("/api/merchant/notice",mc.getNotice)
	engine.GET("/api/merchant/sections",mc.getSections)

}
var(
	ms=service.MerChant{}
)
func(mc *MerchantController)getSections(ctx *gin.Context) {
	typeSlice := []string{"01%", "02%", "03%", "04%", "05%", "06%", "07%", "08%", "09%", "10%", "11%", "13%", "15%"}
	var Date []model.SectionMerchant
	for i, channelType := range typeSlice {
		fmt.Println(i)
		var section model.SectionMerchant
		randSlice, rankSlice, err :=ms.GetMerchantGood(channelType)
		if err!=nil{
			fmt.Println("获取商品路径失败")
			tool.RespInternalError(ctx)
			return
		}
		section.List=randSlice
		section.List=rankSlice

		if section.List==nil{
			section.List=[]model.Merchant{}
		}
		if section.List==nil{
			section.Rank=[]model.Merchant{}
		}
		Date=append(Date,section)

	}
	fmt.Println(Date)
	tool.RespSuccessfulWithDate(ctx,Date)




}

func (mc *MerchantController)getNotice(ctx *gin.Context) {
	//通过商户id获得公告
	merchantNumber:=ctx.Query("merchant_number")
	if merchantNumber==""{
		fmt.Println("商户id不能为空")
		tool.RespInternalError(ctx)
		return
	}
	mId,err:=strconv.ParseInt(merchantNumber,10,64)
	if err!=nil{
		fmt.Println("解析失败")
		tool.RespInternalError(ctx)
		return
	}
	flag,err:=ms.JudgeMerchantId(mId)
	if err!=nil{
		fmt.Println("判断商户id失败")
		tool.RespInternalError(ctx)
		return
	}
	if flag==false{
		tool.RespErrorWithDate(ctx,"商户id错误")
		return
	}
	notice,err:=ms.GetNotice(mId)
	if err!=nil{
		fmt.Println("获取公告失败")
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx,notice)
}



