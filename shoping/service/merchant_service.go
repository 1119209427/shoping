package service

import (
	"shoping/dao"
	"shoping/model"
	"shoping/tool"
)

type MerChant struct {

}
func(mc *MerChant)GetMerchantGood(channel string)([]model.Merchant,[]model.Merchant,error){
	var randSlice, rankSlice []model.Merchant
    md:= dao.MerchantDao {tool.GetDB()}
	var merchantWithGood model.Merchant
	gidSlice,err:=md.QueryRandomGood(channel)
	if err!=nil{
		return nil, nil, err
	}
	for _,gid:=range gidSlice{
		merchant,err:=md.QueryMerchantInfoById(gid)
		if err!=nil{
			return nil, nil, err
		}
		merchantWithGood.Id=merchant.Id
		merchantWithGood.Notice=merchant.Notice
		merchantWithGood.MerchantName=merchant.MerchantName
		merchantWithGood.Description=merchant.Description
		merchantWithGood.Channel=merchant.Channel
		merchantWithGood.Good=merchant.Good
		merchantWithGood.Time=merchant.Time
		merchantWithGood.FavorableRate=merchant.FavorableRate
		merchantWithGood.Volume=merchant.Volume
		merchantWithGood.Followers=merchant.Followers
		merchantWithGood.Followings=merchant.Followings
		randSlice=append(randSlice,merchantWithGood)

	}
	idSlice,err:=md.QueryRankGood(channel)
	if err!=nil{
		return nil, nil, err
	}
	for _,id:=range idSlice{
		merchant,err:=md.QueryMerchantInfoById(id)
		if err!=nil{
			return nil, nil, err
		}
		merchantWithGood.Id=merchant.Id
		merchantWithGood.Notice=merchant.Notice
		merchantWithGood.MerchantName=merchant.MerchantName
		merchantWithGood.Description=merchant.Description
		merchantWithGood.Channel=merchant.Channel
		merchantWithGood.Good=merchant.Good
		merchantWithGood.Time=merchant.Time
		merchantWithGood.FavorableRate=merchant.FavorableRate
		merchantWithGood.Volume=merchant.Volume
		merchantWithGood.Followers=merchant.Followers
		merchantWithGood.Followings=merchant.Followings
		rankSlice=append(rankSlice,merchantWithGood)

	}
	return randSlice,rankSlice,nil




}
func(mc *MerChant)JudgeMerchantId(mid int64)(bool,error){
	_,err:=hd.QueryMerchantInfoById(mid)
	if err!=nil{
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false,err
	}



	return true,nil
}
func(mc *MerChant)GetNotice(mid int64)(string,error){
	merchant,err:=hd.QueryMerchantInfoById(mid)
	if err!=nil{
		return "", err
	}
	return merchant.Notice,nil

}

