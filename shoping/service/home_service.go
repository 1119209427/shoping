package service

import (
	"log"
	"shoping/dao"
	"shoping/model"
	"shoping/tool"
)

type HomeService struct {

}
var (
	hd=dao.MerchantDao{tool.GetDb()}
)

// GetChannelInfo 获取京东快报
func(hc *HomeService)GetChannelInfo()([]model.Good,[]model.Good,error){
	var randSlice, rankSlice []model.Good
	gIdSlice,err:=gd.QueryRandomInfo()
	if err!=nil{
		return nil, nil, err
	}
	for _,gId:=range gIdSlice{
		var good model.Good
		Good,err:=gd.QueryId(gId)
		if err!=nil{
			return nil, nil, err
		}
		good.Id=Good.Id
		good.GoodId=Good.GoodId
		good.BeforePrice=Good.BeforePrice
		good.AfterPrice=Good.AfterPrice
		good.Channel=Good.Channel
		good.Description=Good.Description
		good.Good=Good.Good
		good.Merchant=Good.Merchant
		good.MerchantNumber=Good.MerchantNumber
		good.Time=Good.Time
		good.Comment=Good.Comment
		good.Views=Good.Views
		good.Likes=Good.Likes
		good.FavorableRate=Good.FavorableRate
		good.Volume=Good.Volume
		good.Followers=Good.Followers
		good.Followings=Good.Followings
		randSlice=append(randSlice,good)
	}
	idSlice,err:=gd.QueryRankInfo()
	if err!=nil{
		log.Fatal()
		return nil, nil, err
	}
	for _,id:=range idSlice{
		var good model.Good
		Good,err:=gd.QueryId(id)
		if err!=nil{
			return nil, nil, err
		}
		good.Id=Good.Id
		good.GoodId=Good.GoodId
		good.BeforePrice=Good.BeforePrice
		good.AfterPrice=Good.AfterPrice
		good.Channel=Good.Channel
		good.Description=Good.Description
		good.Good=Good.Good
		good.Merchant=Good.Merchant
		good.MerchantNumber=Good.MerchantNumber
		good.Time=Good.Time
		good.Comment=Good.Comment
		good.Views=Good.Views
		good.Likes=Good.Likes
		good.FavorableRate=Good.FavorableRate
		good.Volume=Good.Volume
		good.Followers=Good.Followers
		good.Followings=Good.Followings
		rankSlice=append(randSlice,good)
	}
	return randSlice,rankSlice,nil


}



// GetChannelGood 通过分区信息获得商品信息
func(hc *HomeService)GetChannelGood(channel string)([]model.Good,[]model.Good,error){
	var randSlice, rankSlice []model.Good
	gIdSlice,err:=gd.QueryRandomChannel(channel)
	if err!=nil{
		return nil, nil, err
	}
	for _,gId :=range gIdSlice{
		var good model.Good
		Good,err:=gd.QueryId(gId)
		if err!=nil{
			return nil, nil, err
		}
		good.Id=Good.Id
		good.GoodId=Good.GoodId
		good.BeforePrice=Good.BeforePrice
		good.AfterPrice=Good.AfterPrice
		good.Channel=Good.Channel
		good.Description=Good.Description
		good.Good=Good.Good
		good.Merchant=Good.Merchant
		good.MerchantNumber=Good.MerchantNumber
		good.Time=Good.Time
		good.Comment=Good.Comment
		good.Views=Good.Views
		good.Likes=Good.Likes
		good.FavorableRate=Good.FavorableRate
		good.Volume=Good.Volume
		good.Followers=Good.Followers
		good.Followings=Good.Followings
		randSlice=append(randSlice,good)
	}
	rankgIdSlice,err:=gd.QueryRankChannel(channel)
	for _,gId:=range rankgIdSlice{
		var good model.Good
		Good,err:=gd.QueryId(gId)

		if err!=nil{
			return nil, nil, err
		}
		good.Id=Good.Id
		good.GoodId=Good.GoodId
		good.BeforePrice=Good.BeforePrice
		good.AfterPrice=Good.AfterPrice
		good.Channel=Good.Channel
		good.Description=Good.Description
		good.Good=Good.Good
		good.Merchant=Good.Merchant
		good.MerchantNumber=Good.MerchantNumber
		good.Time=Good.Time
		good.Comment=Good.Comment
		good.Views=Good.Views
		good.Likes=Good.Likes
		good.FavorableRate=Good.FavorableRate
		good.Volume=Good.Volume
		good.Followers=Good.Followers
		good.Followings=Good.Followings
		rankSlice=append(rankSlice,good)
	}
	return randSlice,rankSlice,nil





}

func(hc *HomeService)Search(keyWord string)([]model.MerchantWithGood,error){
	var result []model.MerchantWithGood
	var merchantWithGood model.MerchantWithGood
	gdSlice,err:=gd.SearchIdByKeywords(keyWord)
	if err!=nil{
		return nil, err
	}
	for _,goodid:=range gdSlice{
		//获取商品信息，通过商品id
		goodModel,err:=gd.QueryId(goodid)
		if err!=nil{
			return nil, err
		}
		//通过商品信息获得店铺信息
		merchant,err:=hd.QueryMerchantInfoById(goodModel.MerchantNumber)
		if err!=nil{
			return nil, err
		}
		merchantWithGood.Merchant=merchant
		merchantWithGood.Id=merchant.Id
		merchantWithGood.MerchantName=merchant.MerchantName
		merchantWithGood.Description=merchant.Description
		merchantWithGood.Channel=merchant.Channel
		merchantWithGood.Time=merchant.Time
		//merchantWithGood.FavorableRate=merchant.FavorableRate
		merchantWithGood.Volume=merchant.Volume
		merchantWithGood.Followers=merchant.Followers
		merchantWithGood.Followings=merchant.Followings
		result=append(result,merchantWithGood)

	}
	return result,nil





}