package model
type ShoppingCart struct {
	Id int64
	CartId int64
	CartItems []*CartItem//所有的商品
	TotalCount int64 //商品数总和
	TotalAmount float64//金额总和
	UserId int64 //购物车所属的用户
}

// CartItem 购物车里的一项商品
type CartItem struct {
	Good *Good
	Count int64//购物车中一项商品的数目
	CartId int64 //属于哪一个购物车
	Amount float64 //价格



}
