package model

import "time"

type User struct {
	Id int64
	TotalLikes      int64
	Username        string
	Password        string
	Email           string
	Phone           string
	Salt            string
	Avatar          string
	RegDate         time.Time
	Statement       string
	Gender          string
	Balance         float64//余额
	CartId          int64//购物车id
	Token      string





}
