package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/smtp"
	"shoping/dao"
	"shoping/model"
	"shoping/tool"
	"strconv"
	"strings"
	"time"
)

type UserService struct {

}
var(
	ud=dao.UserDao{tool.GetDB()}
	rd=dao.RedisDao{}
)
func(us *UserService)ChangeBalance(change ,uid int64 )error{
	err:=ud.UpdateBalance(change,uid)
	if err!=nil{
		return err
	}
	return nil


}
func(us *UserService)JudgeUid(id int64)(bool,error){
	_,err:=ud.QueryUserInfoById(id)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			return false, nil
		}
		return false, err
	}
	return true,nil
}
func(us *UserService)GetUserInfo(id int64)(model.User,error){
	//获取基本信息
	user,err:=ud.QueryUserInfoById(id)
	if err!=nil{
		return user,err
	}
	return user,nil

}

// ChangePwd 改变密码
func(us *UserService)ChangePwd(id int64,newPwd string)error{
	//加盐
	salt:=strconv.FormatInt(time.Now().Unix(),10)
	hash:=md5.New()
	hash.Write([]byte(newPwd))
	hash.Write([]byte(salt))
	st:=hash.Sum(nil)
	saltedPassword :=hex.EncodeToString(st)
	_,err:=ud.UpdatePwd(id,salt,saltedPassword)
	if err!=nil{
		return err
	}
	return nil


}

// ChangeUserName 改变用户名
func(us *UserService)ChangeUserName(id int64,newUserName string)error{
	err:=ud.ChangeUserNameById(id,newUserName)
	if err!=nil{
		return err
	}
	return nil


}

// ChangeGenderById 改变性别
func(us *UserService)ChangeGenderById(id int64,newGender string)error{
	err:=ud.ChangeGenderById(id,newGender)
	if err!=nil{
		return err
	}
	return nil
}

// ChangePhoneById 改变手机号
func(us *UserService)ChangePhoneById(id int64,phone string)error{
	err:=ud.ChangePhoneById(id,phone)
	if err!=nil{
		return err
	}
	return nil
}

// IsRepeatPhone 检查手机号是否重复
func(us *UserService)IsRepeatPhone(phone string)(bool,error){
	_,err:=ud.QueryByPhone(phone)
	if err!=nil{
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false,err
	}
	return true ,nil


}

// ChangeEmail 修改邮箱
func (us *UserService)ChangeEmail(id int64,newEmail string)error{
	newEmail=strings.ToLower(newEmail)
	err:=ud.ChangeEmailById(id,newEmail)
	if err!=nil{
		return err
	}
	return nil

}

// IsRepeatEmail 验证邮箱是否重复
func(us *UserService)IsRepeatEmail(email string)(bool,error){
	email=strings.ToLower(email)
	_,err:=ud.QueryByEmail(email)
	if err!=nil{
		log.Fatal(err.Error())
		return false,err
	}

	return true,nil

}

// VerifyCodeIn 将验证码放入redis
func (us *UserService) VerifyCodeIn(ctx *gin.Context, key string, value string)error{
	err:=rd.RedisSetValue(ctx,key,value)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil

}
// IsRepeatName 检查是否为重复验证码
func (us *UserService) IsRepeatName(username string)(bool,error){
	_,err:=ud.QueryByUsername(username)
	if err!=nil{
		return false,err
	}

return true,nil
}

// SendEmailCode 通过邮箱发送验证码
func(us *UserService)SendEmailCode(email string)(string,error){
	email=strings.ToLower(email)
	emailCfg:=tool.GetCfg().Email
	to:=[]string{email}
	fmt.Println("EMAIL", email)
	auto:=smtp.PlainAuth("",emailCfg.ServiceEmail,emailCfg.ServicePwd,emailCfg.SmtpHost)
	code:=fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	str:=fmt.Sprintf("From:%v\r\nTo:%v\r\nSubject:jindong验证码\r\n\r\n您的验证码为：%s\r\n请在10分钟内完成验证", emailCfg.ServiceEmail, email, code)
	msg:=[]byte(str)
	err:=smtp.SendMail(emailCfg.SmtpHost+":"+emailCfg.SmtpPort,auto,emailCfg.ServiceEmail,to,msg)
	if err!=nil{
		log.Fatal(err.Error())
		return "",err
	}
	return code,nil

}

// SelectIdByEmail 通过邮箱地址获得id
func(us *UserService)SelectIdByEmail(email string)(int64,error){
	user,err:=ud.SelectIdByEmail(email)
	if err!=nil{
		return 0,err
	}
	return user.Id,err

}

// SelectIdByPhone 通过手机号获得id
func(us *UserService)SelectIdByPhone(phone string)( int64,error){
	user,err:=ud.SelectIdByPhone(phone)
	if err!=nil{
		return 0,err

	}
	return user.Id,err

}

// JudgePhone 判断手机号是否存在
func(us *UserService)JudgePhone(phone string)(bool,error){
	//查询数据库

	_,err:=ud.QueryByPhone(phone)
	if err!=nil{
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false,err
	}
	return true,nil
}

// JudgeVerifyCode 判断验证码是否真确
func(us *UserService)JudgeVerifyCode(ctx *gin.Context, key string, givenValue string) (bool, error){
	rd:=dao.RedisDao{}
	value,err:=rd.RedisGetValue(ctx,key)
	if err!=nil{

		log.Fatal(err.Error())
		return false,err
	}
	if value!=givenValue{
		return false,nil
	}
	return true,nil

}

// InsertUser 将注册成功的用户放入
func(us *UserService)InsertUser(user model.User)error{
	err:=ud.InsertUser(user)
	if err!=nil{
		return err
	}
	return nil

}

// LoginByPwd 通过密码登录，返回一个实体
func(us *UserService)LoginByPwd(logins,password string)(model.User,bool,error){
	//判断登录类型
	index:=strings.Index(logins,"@")
	if index!=-1{
		loginname:=strings.ToLower(logins)
		var user, err = ud.QueryByEmail(loginname)
		if err!=nil{
			if err.Error()=="sql: no rows in result set" {
				return model.User{}, false, nil
			}
			return model.User{}, false, err
		}
		//使用md5解密
		hash:=md5.New()
		hash.Write([]byte(password))
		hash.Write([]byte(user.Salt))
		st:=hash.Sum(nil)
		hashPwd := hex.EncodeToString(st)

		if hashPwd != user.Password {
			return model.User{}, false, nil
		}
		return user, true, nil

	}else{
		//手机登录
		user,err:=ud.QueryByPhone(logins)
		if err!=nil{
			if err.Error()=="sql: no rows in result set"{
				return model.User{}, false, nil
			}
			return model.User{}, false, err
			//使用md5解析
			hash:=md5.New()
			hash.Write([]byte(password))
			hash.Write([]byte(user.Salt))
			st:=hash.Sum(nil)
			hashPwd:=hex.EncodeToString(st)
			if hashPwd!=user.Password{
				return model.User{}, false, err
			}

		}
		return user, true, nil

	}


}