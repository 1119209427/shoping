package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"shoping/dao"
	"shoping/model"
	"shoping/tool"
	"strconv"
)

type GoodService struct {

}
var(
	gd=dao.GoodDao{tool.GetDb()}

)

// GoodInShoppingCart 根据商品id把商品保存入购物车
func(gs *GoodService)GoodInShoppingCart(ctx *gin.Context,key int64,value model.Good)error{
	keyString:=strconv.FormatInt(key,10)
	err:=rd.RedisSetValueShop(ctx,keyString,value)
	if err!=nil{
		return err
	}
	return nil


}

// GetLikesStatus 获取用户喜爱状态，在err为nil的情况下，已经喜爱返回true，反之返回false
//考虑到单个商品可能存在大量喜爱，这里在dao层查询用户喜爱的商品，而不是查询喜爱过商品的用户，优化性能
func(gs *GoodService)GetLikesStatus(Gid,Uid int64)(bool,error){
	GidSlice,err:=gd.QueryByGid(Uid)
	if err!=nil{
		log.Fatal(err.Error())
		return false, err
	}
	for _,v:=range GidSlice{
		if v==Gid{
			return true,nil
		}
	}
	return false, err
}
// IncreaseLikes 增加点赞数
func(gs *GoodService)IncreaseLikes(
	flag bool,
	uid,Gid int64,
)error{
	//若没有点赞
	if flag==false{
		//将数据插入like表
		err:=gd.InsertLike(uid,Gid)
		if err!=nil{
			return err
		}
		//更新商品的喜爱数量
		err=gd.UpdateLikes(Gid,1)
		if err!=nil{
			return err
		}
		//改变用户的喜爱数量
		err=gd.UpdateTotalLikes(uid,1)
		if err!=nil{
			return err
		}

	}else{
		//若点赞了
		err:=gd.DeleteLikes(uid,Gid)
		if err!=nil{
			return err
		}
		err=gd.UpdateLikes(Gid,-1)
		if err!=nil{
			return err
		}
		err=gd.UpdateTotalLikes(uid,-1)
		if err!=nil{
			return err
		}


	}
	return nil

}

// SolveFollow 更新关注状态
func(gs *GoodService)SolveFollow(flag bool, followerUid int64, followingGood int64) error{
	//若没有关注
	if flag==false{
		//1将id插入
		err:=gd.InsertFollow(followerUid,followingGood)
		if err!=nil{
			return err
		}
		//2.被关注用户更新关注者数量
		err=gd.UpdateFollower(followingGood,1)
		if err!=nil{
			return err

		}
		//3更新关注中数量
		err=gd.UpdateFollowing(followingGood,1)
		if err!=nil{
			return err
		}
	}else{
		err:=gd.DeleteFollow(followerUid,followingGood)
		if err!=nil{
			return err
		}
		//2.被关注用户更新关注者数量
		err=gd.UpdateFollower(followingGood,-1)
		if err!=nil{
			return err

		}
		//3更新关注中数量
		err=gd.UpdateFollowing(followingGood,-1)
		if err!=nil{
			return err
		}


	}
	return nil




}


// JudgeGoodId 判断商品id
func(gs *GoodService)JudgeGoodId(id int64)(bool,error) {
	_, err := gd.QueryId(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, err
		}
		return false, err
	}
	return true, err
}

// GetGoodInfo 获取商品信息
func(gs *GoodService)GetGoodInfo(id int64)(model.Good,error){
	good,err:=gd.QueryId(id)
	if err!=nil{
		return good,err
	}
	return good,nil

}

// GetFollowStatus 获取关注状态
func(gs *GoodService)GetFollowStatus(followingId, followedId int64)(bool,error){
	id,err:=gd.QueryFollowedId(followingId)
	if err!=nil{
		return false,err
	}
	for _, userid :=range id{
		if userid==followedId{
			return true,nil

		}


	}


	return false,nil







}