package param

import "shoping/model"

type CommentParam struct {
	Id      int64
	User    model.User
	GoodId int64
	Value   string
	Time    string
	Likes   int64
}

