package tool

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"shoping/model"
)

var DB *gorm.DB
func InitGorm()error{
	host:="127.0.0.1"
	port:="3306"
	username:="root"
	password:="123456"
	dbname:="shopping"
	dsn := username+":"+password+"@("+host+":"+port+")/"+dbname+"?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open(mysql.Open(dsn) )
	if err!=nil{
		log.Fatal("init the mysql failed",err.Error())
		return err
	}
	DB=db
	tableInit()
	return nil



	/*db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		Conn:                      nil,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         171,
		DefaultDatetimePrecision:  nil,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
		DontSupportForShareClause: false,

	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating:true,//逻辑外键（代码里面自动外键外键关系)
	})
	if err!=nil{
		panic(err)
	}
	DB=db
	sqlDB, err := db.DB()
	if err!=nil{
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db*/


}
func tableInit(){
	var user model.User
	DB.AutoMigrate(&user)
	var comment model.Comment
	DB.AutoMigrate(&comment)
	var good model.Good
	var goodlike model.Good
	DB.AutoMigrate(&good,&goodlike)

}
func GetGormDb() *gorm.DB{
	return DB
}