package tool

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var dB *sql.DB
func GetDB()*sql.DB{
	return dB
}
func InitDB (){
	cfg:=GetCfg().DataBase
	db,err:=sql.Open(cfg.Driver, cfg.User+":"+cfg.Password+"@tcp("+cfg.Host+":"+cfg.Port+")/"+cfg.DbName+"?charset=utf8&parseTime=true&loc=Local")
	if err!=nil{
		panic(err)
	}
	dB=db

}
