package dao

import (
	"database/sql"
	"log"
	"shoping/model"
)

type GoodDao struct {
	*sql.DB
}

// QueryRankInfo 按照点击率排行4个，返回商品id
func(gd *GoodDao) QueryRankInfo()([]int64,error) {
	var goodidSlice []int64
	stmt, err := gd.Prepare("select goodid from good where goodid in(select goodid from good_label where good_label =?) order by views desc limit 4")
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query("京东快报")
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	for rows.Next(){
		var goodid int64
		err=rows.Scan(&goodid)
		if err!=nil{
			log.Fatal(err.Error())
			return nil, err
		}
		goodidSlice=append(goodidSlice,goodid)

	}
	return goodidSlice,nil

}






// QueryRandomInfo 京东快报，返回商品id
func(gd *GoodDao)QueryRandomInfo()([]int64,error){
	var goodIdSlice []int64
	stmt,err:=gd.Prepare("select goodid from good_label where good_label = ? order by rand() limit 10")
	if err!=nil{
		log.Fatal(err.Error())
		return nil, err
	}
	defer stmt.Close()
	rows,err:=stmt.Query("京东快报")
	if err!=nil{
		log.Fatal(err.Error())
		return nil, err
	}
	for rows.Next(){
		var goodid int64

		err:=rows.Scan(&goodid)
		if err!=nil{
			log.Fatal(err.Error())
			return nil, err
		}
		goodIdSlice=append(goodIdSlice,goodid)


	}
	return goodIdSlice,nil
}

// QueryRankChannel 根据成交量排行返回特定分区的商品15个，返回商品id
func(gd *GoodDao)QueryRankChannel(channel string)([]int64,error){
	var goodIdSlice []int64
	stmt,err:=gd.Prepare("select goodid from good where channel like ? order by views desc limit 15")
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
		var goodid int64
		rows.Scan(&goodid)
		if err!=nil{
			log.Fatal(err.Error())
			return nil, err
		}
		goodIdSlice=append(goodIdSlice,goodid)

	}
	return goodIdSlice,nil

}
// QueryRandomChannel 通过分区的商品，返回商品id
func(gd *GoodDao)QueryRandomChannel(channel string)([]int64,error){
	var goodIdSlice []int64
	stmt,err:=gd.Prepare("select id,goodid,before_price,after_price,channel,description,good,merchant,merchant_number,time,comment,views,likes,favorable_rate,volume,followers,followings from good where channel = ?")
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
	var good model.Good
	for rows.Next(){
		err:=rows.Scan(&good.Id,&good.GoodId,&good.BeforePrice,&good.AfterPrice,&good.Channel,&good.Description,&good.Good,&good.Merchant,&good.MerchantNumber,&good.Time,&good.Comment,&good.Views,&good.Likes,&good.FavorableRate,&good.Volume,&good.Followers,&good.Followings)
		if err!=nil{
			log.Fatal(err.Error())
			return nil, err
		}
		var goodid int64
		goodIdSlice=append(goodIdSlice,goodid)
	}
	return goodIdSlice,nil


}
func(gd *GoodDao)SearchIdByKeywords(keywords string)([]int64,error){
	var gdSlice []int64
	keywords = "%" + keywords + "%"
	stmt,err:=gd.Prepare("select goodid from good where title like ? order by views desc ")
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}
	defer stmt.Close()
	rows,err:=stmt.Query(keywords)
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}
	for rows.Next(){
		var goodid int64
		err=rows.Scan(&goodid)
		if err != nil {
			return nil, err
		}
		gdSlice=append(gdSlice,goodid)

	}
	return gdSlice,nil
}
func(gd *GoodDao)DeleteLikes(uid,gid int64)error{
	stmt,err:=gd.Prepare("delete from good_like where (uid,goodid) values (?,?)")
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(uid,gid)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil



}

// UpdateTotalLikes 增加用户的喜爱数
func(gd *GoodDao)UpdateTotalLikes(uid,num int64)error{
	stmt,err:=gd.Prepare("update user set total_likes=total_likes+? where id=?")
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(num,uid)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil

}


// UpdateLikes 增加商品喜爱数
func(gd *GoodDao)UpdateLikes(gid,num int64)error{
	stmt,err:=gd.Prepare("update good set likes=likes+? where goodid = ?")

	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(num,gid)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil
}
// InsertLike 插入喜爱表
func(gd *GoodDao)InsertLike(uid,gid int64)error{
	stmt,err:=gd.Prepare("insert into good_like (uid,goodid) values (?,?)")
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(uid,gid)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil
}

// QueryByGid  查询用户喜爱的商品id
func(gd *GoodDao)QueryByGid(uid int64)([]int64,error){
	var GidSlice []int64
	stmt,err:=gd.Prepare("select goodid from good_like where uid =? ")
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}
	defer stmt.Close()
	rows,err:=stmt.Query(uid)
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}
	for rows.Next(){
		var goodid int64
		err=rows.Scan(&goodid)
		if err!=nil{
			log.Fatal(err.Error())
			return nil, err
		}
		GidSlice=append(GidSlice,goodid)
	}
	return GidSlice,nil
}

func(gd *GoodDao)DeleteFollow(followingId,followingGoodId int64)error{
	stmt,err:=gd.Prepare("delete from follow_good where(following_goodid=? and follower_id=? )")
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(followingId,followingGoodId)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}

	return nil
}
func(gd *GoodDao)UpdateFollowing(followingId,num int64)error{
	stmt,err:=gd.Prepare("update good set following=following + ? where goodid = ?")
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(num,followingId)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil

}
func(gd *GoodDao)UpdateFollower(followingGoodId,num int64)error{
	stmt,err:=gd.Prepare("update good set followers=followers + ? where goodid =?")
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(num,followingGoodId)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil

}
func(gd *GoodDao)InsertFollow(followingId,followingGood int64)error {
	stmt,err:=gd.Prepare("insert into good_follow(following_id,follower_id) values (?,?)")

	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(followingId,followingGood)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil


}
func (gd *GoodDao)QueryId(id int64)(model.Good,error) {
	good:=model.Good{}
	stmt,err:=gd.Prepare("select id,before_price,after_price,channel,description, good merchant, merchantnumber, time, comment, likes, favorablerate,volume  from good where id = ?")
	if err!=nil{
		log.Fatal(err.Error())
		return good,err
	}
	defer stmt.Close()
	row:=stmt.QueryRow(id)
	if row.Err()!=nil{
		log.Fatal(row.Err())
		return good,row.Err()
	}
	err=row.Scan(&good.Id,&good.BeforePrice,&good.AfterPrice,&good.Channel,&good.Description,&good.Merchant,&good.MerchantNumber,&good.Time,&good.Comment,&good.Likes,&good.FavorableRate,&good.Volume)
	if err!=nil{
		log.Fatal(err.Error())
		return good,err
	}
	return good,nil

}
func(gd *GoodDao)QueryFollowedId(followingId int64)([]int64,error){
	var id []int64
	stmt,err:=gd.Prepare("select following_id from good_follow where follower_id=?")
	if err!=nil{
		log.Fatal(err.Error())
		return nil, err
	}
	defer stmt.Close()
	rows,err:=stmt.Query(followingId)
	if err!=nil{
		log.Fatal(err.Error())
		return nil, err
	}
	for rows.Next(){
		var goodId int64
		rows.Scan(&goodId)
		if err!=nil{
			log.Fatal(err.Error())
			return nil, err

		}
		id=append(id,goodId)
	}
	return id,nil


}
