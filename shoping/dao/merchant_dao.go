package dao

import (
	"database/sql"
	"log"
	"shoping/model"
)

type MerchantDao struct {
	*sql.DB

}
func(md *MerchantDao)QueryRankGood(channel string)([]int64,error){
	var gidSlice []int64
	stmt,err:=md.Prepare("select goodid from merchant where channle like ? order by views desc limit 10")
	if err!=nil{
		log.Fatal(err.Error())
		return nil, err
	}
	defer stmt.Close()
	rows,err:=stmt.Query(channel)
	if err!=nil{
		log.Fatal(err.Error())
		return nil, err
	}
	for rows.Next(){
		var gid int64
		err=rows.Scan(&gid)
		if err!=nil{
			log.Fatal(err.Error())
			return nil, err
		}
		gidSlice=append(gidSlice,gid)
	}
	return gidSlice,nil

}

// QueryRandomGood 返回商户的商品id
func(md *MerchantDao)QueryRandomGood(channel string)([]int64,error){
	var gidSlice []int64
	stmt,err:=md.Prepare("select goodid from merchant where channle like ? order by rand()")
	if err!=nil{
		log.Fatal(err.Error())
		return nil, err
	}
	defer stmt.Close()
	rows,err:=stmt.Query(channel)
	if err!=nil{
		log.Fatal(err.Error())
		return nil, err
	}
	for rows.Next(){
		var gid int64
		err=rows.Scan(&gid)
		if err!=nil{
			log.Fatal(err.Error())
			return nil, err
		}
		gidSlice=append(gidSlice,gid)
	}
	return gidSlice,nil
}
// QueryMerchantInfoById 通过merchant_id获得店家的详细
func(md *MerchantDao)QueryMerchantInfoById(id int64)(model.Merchant,error){
	merchant:=model.Merchant{}
	stmt,err:=md.Prepare("select id,merchant_number,merchant_name,description,channel,good,time,favorable_rate,volume,followers,followings from merchant where merchant_number = ? ")
	if err!=nil{
		log.Fatal(err.Error())
		return merchant,err
	}
	defer stmt.Close()
	row:=stmt.QueryRow(id)
	if row.Err()!=nil{
		log.Fatal(err.Error())
		return model.Merchant{}, err
	}


	err=row.Scan(&merchant.Id,&merchant.MerchantNumber,&merchant.MerchantName,&merchant.Description,&merchant.Channel,&merchant.Good,&merchant.Time,&merchant.FavorableRate,&merchant.Volume,&merchant.Followers,&merchant.Followings)
	if err!=nil{
		log.Fatal(err.Error())
		return merchant, err

	}
	return merchant, nil


}
