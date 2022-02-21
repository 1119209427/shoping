package dao

import (
	"database/sql"
	"log"
	"shoping/model"
)

type CartDao struct {
	*sql.DB
}
func(c *CartDao)QueryByCid(cartId int64)(model.ShoppingCart,error){
	cart:=model.ShoppingCart{}
	stmt,err:=c.Prepare("select id,cart_id,cart_items,total_count,total_amount,user_id from shoppingcart where cart_id = ?")
	if err!=nil{
		log.Fatal(err.Error())
		return cart,err
	}
	defer stmt.Close()
	row:=stmt.QueryRow(cartId)
	if row.Err()!=nil{
		log.Fatal(err.Error())
		return cart,row.Err()
	}
	err=row.Scan(&cart.Id,&cart.CartId,&cart.CartItems,&cart.TotalCount,&cart.TotalAmount,&cart.UserId)
	if err!=nil{
		log.Fatal(err.Error())
		return cart,err
	}
	return cart,nil
}
