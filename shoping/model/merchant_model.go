package model

import "time"

type Merchant struct {
	Id int64
	Notice string
	MerchantNumber   int64//商家id
	MerchantName string
	Description string
	Channel     string
	Good        []Good
	Time        time.Time
	FavorableRate     float64//好评率
	Volume      int64//成交量
	Followers       int64
	Followings      int64
}