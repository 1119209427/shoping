package model

import (
	"gorm.io/gorm"
	"time"
)

type Good struct {
	gorm.Model
	Id             int64
	GoodId         int64
	BeforePrice    float64
	AfterPrice     float64
	Channel        string
	Description    string
	Good           string
	Merchant       string
	MerchantNumber int64 //商家id
	Time           time.Time
	Comment_number int64
	Views          int64
	Likes          int64
	FavorableRate  int64 //好评率
	Volume         int64 //成交量
	Followers      int64
	Followings     int64
	// Good拥有并属于多种 FollowUser，`good_follow_user` 是连接表
	FollowUser []User `gorm:"many2many:good_follow_user"`
	Comment    []Comment
}
type GoodFollow struct {
	//多对多（商品为一，顾客关注为多)
	FollowingId []int64 //顾客id
	FollowerId  int64   //商品id
	Good        Good
}
type GoodLike struct {
	gorm.Model
	Uid    int64
	GoodId int64
}
type MerchantWithGood struct {
	Merchant       Merchant
	Id             int64
	BeforePrice    float64
	AfterPrice     float64
	Channel        string
	Description    string
	Good           string
	MerchantName   string
	MerchantNumber int64
	Time           time.Time
	Comment        int64
	Views          int64
	Likes          int64
	FavorableRate  int64 //好评率
	Volume         int64 //成交量
	Followers      int64
	Followings     int64
}
