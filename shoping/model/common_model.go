package model

import "time"

type Comment struct {
	Id      int64
	GoodId int64
	UserId  int64
	Value   string
	Time    time.Time
	Likes   int64
}
