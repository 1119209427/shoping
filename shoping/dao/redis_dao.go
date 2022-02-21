package dao

import (
	"github.com/gin-gonic/gin"
	"shoping/model"
	"shoping/tool"
	"time"
)

type RedisDao struct{

}
func (rd *RedisDao)RedisGetValue(ctx *gin.Context,key string)(string,error){
	redisConn:=tool.GetRedisConn()
	stringcmd:=redisConn.Get(ctx,key)
	if stringcmd.Err()!=nil{
		return "",stringcmd.Err()
	}
	return stringcmd.Val(),nil
}
func(rd *RedisDao)RedisSetValue(ctx *gin.Context,key string,value string)error{
	redisConn:=tool.GetRedisConn()
	statusCmd:=redisConn.Set(ctx,key,value,time.Minute*5)
	return statusCmd.Err()

}
func(rd *RedisDao)RedisGetValueShop(ctx *gin.Context,key string)(model.Good,error){
	good:=model.Good{}
	redisConn:=tool.GetRedisConn()
	stringcmd:=redisConn.Get(ctx,key)
	if stringcmd.Err()!=nil{
		return good,stringcmd.Err()
	}
	return good,nil
}
func(rd *RedisDao)RedisSetValueShop(ctx *gin.Context,key string,value model.Good)error{
redisConn:=tool.GetRedisConn()
statusCmd:=redisConn.Set(ctx,key,value,time.Hour*24*7)//保存七天
return statusCmd.Err()

}
