package model
type Order struct {
	Type string //待付款，待发货，待收货
	Uid int64
	GoodId int64
	Good Good
}
