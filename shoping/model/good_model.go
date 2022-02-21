package model

import "time"

type Good struct {
	Id          int64
	GoodId      int64
	BeforePrice float64
	AfterPrice  float64
	Channel     string
	Description string
	Good        string
	Merchant    string
	MerchantNumber   int64//商家id
	Time           time.Time
	Comment_number int64
	Views          int64
	Likes        int64
	FavorableRate     int64//好评率
	Volume      int64//成交量
	Followers       int64
	Followings      int64
	Comment  Comment
}
type GoodFollow struct {
	FollowingId int64
	FollowerId  int64
	Good Good
}
type MerchantWithGood struct{
	Merchant  Merchant
	Id          int64
	BeforePrice float64
	AfterPrice  float64
	Channel     string
	Description string
	Good        string
	MerchantName   string
	MerchantNumber   int64
	Time        time.Time
	Comment      int64
	Views       int64
	Likes        int64
	FavorableRate     int64//好评率
	Volume      int64//成交量
	Followers       int64
	Followings      int64

}
