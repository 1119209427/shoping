package dao

import (
	"database/sql"
	"log"
)

type OrderDao struct {
	*sql.DB
}
func (od *OrderDao)GetGoodUid(uid int64,state string)([]int64,error){
	var goodIdSlice []int64

	stmt,err:=od.Prepare("select good_id from order where ( uid = ? and type = ? )")
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}
	defer stmt.Close()
	rows,err:=stmt.Query(uid,state)
	for rows.Next(){
		var uid int64
		err:=rows.Scan(&uid)
		if err!=nil{
			log.Fatal(err.Error())
			return nil,err
		}
		goodIdSlice=append(goodIdSlice,uid)
	}
	return goodIdSlice,nil
}
func (od *OrderDao)ChangeState(uid int64,newType string)error{

	stmt,err:=od.Prepare("update order set type = ? where uid =?")
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(newType,uid)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil
}
