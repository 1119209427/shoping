package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

var (
	ud = dao.UserDao{tool.GetDb()}
	rd = dao.RedisDao{}
)

func (us *UserService) ChangeBalance(change, uid int64) error {
	db := tool.GetGormDb()
	if err := db.Model(&model.User{}).Update("balance", gorm.Expr("balance + ?", change)).Where("id", uid).Error; err != nil {
		return err
	}
	/*err:=ud.UpdateBalance(change,uid)
	if err!=nil{
		return err
	}*/
	return nil

}
func (us *UserService) JudgeUid(id int64) (bool, error) {
	/*_,err:=ud.QueryUserInfoById(id)*/
	db := tool.GetGormDb()
	var user model.User
	var count int64
	if err := db.Model(&model.User{}).Where("id", id).First(&user).Count(&count).Error; err != nil {
		return false, err
		if err.Error() == "record not found" {

			return false, nil
		}
	}
	if count == 1 { //说明找到了
		return true, nil

	}
	/*if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			return false, nil
		}
		return false, err
	}
	return true,nil*/
	return false, nil
}
func (us *UserService) GetBalance(id int64) (float64, error) {
	var user model.User
	db := tool.GetGormDb()
	if err := db.Model(&model.User{}).Where("id", id).First(&user).Error; err != nil {
		return user.Balance, err
	}
	return user.Balance, nil

}

func (us *UserService) GetUserInfo(id int64) (model.User, error) {
	//获取基本信息

	user, err := ud.QueryUserInfoById(id)
	if err != nil {
		return user, err
	}
	return user, nil

}

// ChangePwd 改变密码
func (us *UserService) ChangePwd(id int64, newPwd string) error {
	//加盐
	salt := strconv.FormatInt(time.Now().Unix(), 10)
	hash := md5.New()
	hash.Write([]byte(newPwd))
	hash.Write([]byte(salt))
	st := hash.Sum(nil)
	saltedPassword := hex.EncodeToString(st)
	db := tool.GetGormDb()
	if err := db.Model(&model.User{}).Where("id", id).Update("password", saltedPassword).Error; err != nil {
		return err

	}
	//_,err:=ud.UpdatePwd(id,salt,saltedPassword)
	//if err!=nil{
	//	return err
	//}
	return nil

}

// ChangeUserName 改变用户名
func (us *UserService) ChangeUserName(id int64, newUserName string) error {
	err := ud.ChangeUserNameById(id, newUserName)
	if err != nil {
		return err
	}
	return nil

}

// ChangeGenderById 改变性别
func (us *UserService) ChangeGenderById(id int64, newGender string) error {
	db := tool.GetGormDb()
	if err := db.Model(&model.User{}).Where("id", id).Update("gender", newGender).Error; err != nil {
		return err
	}
	/*	err:=ud.ChangeGenderById(id,newGender)
		if err!=nil{
			return err
		}*/
	return nil
}

// ChangePhoneById 改变手机号
func (us *UserService) ChangePhoneById(id int64, phone string) error {
	db := tool.GetGormDb()
	if err := db.Model(&model.User{}).Where("id", id).Update("phone", phone).Error; err != nil {
		return err
	}
	/*err:=ud.ChangePhoneById(id,phone)
	if err!=nil{
		return err
	}*/
	return nil
}

// IsRepeatPhone 检查手机号是否重复
func (us *UserService) IsRepeatPhone(phone string) (bool, error) {
	var user model.User
	var count int64
	db := tool.GetGormDb()
	if err := db.Model(&model.User{}).Where("phone", phone).First(&user).Count(&count).Error; err != nil {
		if err.Error() == "record not found" {
			return false, nil

		}
		return false, err
	}
	if count == 1 {
		return true, nil
	}

	/*_,err:=ud.QueryByEmail(email)
	if err!=nil{
		log.Fatal(err.Error())
		return false,err
	}*/

	return false, nil

	/*_,err:=ud.QueryByPhone(phone)
	if err!=nil{
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false,err
	}
	return true ,nil*/

}

// ChangeEmail 修改邮箱
func (us *UserService) ChangeEmail(id int64, newEmail string) error {
	newEmail = strings.ToLower(newEmail)
	db := tool.GetGormDb()
	if err := db.Model(&model.User{}).Where("id", id).Update("email", newEmail).Error; err != nil {
		return err
	}

	/*err:=ud.ChangeEmailById(id,newEmail)
	if err!=nil{
		return err
	}*/
	return nil

}

// IsRepeatEmail 验证邮箱是否重复
func (us *UserService) IsRepeatEmail(email string) (bool, error) {
	email = strings.ToLower(email)
	var user model.User
	var count int64
	db := tool.GetGormDb()
	if err := db.Model(&model.User{}).Where("email", email).First(&user).Count(&count).Error; err != nil {
		if err.Error() == "record not found" {
			return false, nil

		}
		return false, err
	}
	if count == 1 {
		return true, nil
	}

	/*_,err:=ud.QueryByEmail(email)
	if err!=nil{
		log.Fatal(err.Error())
		return false,err
	}*/

	return false, nil

}

// VerifyCodeIn 将验证码放入redis
func (us *UserService) VerifyCodeIn(ctx *gin.Context, key string, value string) error {
	err := rd.RedisSetValue(ctx, key, value)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil

}

// IsRepeatName 检查是否为重复用户名
func (us *UserService) IsRepeatName(username string) (bool, error) {
	db := tool.GetGormDb()
	var user model.User
	var count int64
	if err := db.Model(&model.User{}).Where("username", username).First(&user).Count(&count).Error; err != nil {
		if err.Error() == "record not found" {
			return false, nil

		}
		return false, err
	}
	if count == 1 {
		return true, nil
	}
	/*_,err:=ud.QueryByUsername(username)
	if err!=nil{
		return false,err
	}*/

	return false, nil
}

// SendEmailCode 通过邮箱发送验证码
func (us *UserService) SendEmailCode(email string) (string, error) {
	email = strings.ToLower(email)
	emailCfg := tool.GetCfg().Email
	to := []string{email}
	fmt.Println("EMAIL", email)
	auto := smtp.PlainAuth("", emailCfg.ServiceEmail, emailCfg.ServicePwd, emailCfg.SmtpHost)
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	str := fmt.Sprintf("From:%v\r\nTo:%v\r\nSubject:jindong验证码\r\n\r\n您的验证码为：%s\r\n请在10分钟内完成验证", emailCfg.ServiceEmail, email, code)
	msg := []byte(str)
	err := smtp.SendMail(emailCfg.SmtpHost+":"+emailCfg.SmtpPort, auto, emailCfg.ServiceEmail, to, msg)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return code, nil

}

// SelectIdByEmail 通过邮箱地址获得id
func (us *UserService) SelectIdByEmail(email string) (int64, error) {
	var user model.User
	db := tool.GetGormDb()
	if err := db.Model(&model.User{}).Where("email", email).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return 0, nil
		}
		return 0, err
	}
	/*user,err:=ud.SelectIdByEmail(email)
	if err!=nil{
		return 0,err
	}*/
	return user.Id, nil

}

// SelectIdByPhone 通过手机号获得id
func (us *UserService) SelectIdByPhone(phone string) (int64, error) {
	db := tool.GetGormDb()
	var user model.User
	if err := db.Model(&model.User{}).Where("phone", phone).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return 0, nil
		}
		return 0, err
	}
	/*user,err:=ud.SelectIdByPhone(phone)
	if err!=nil{
		return 0,err

	}*/
	return user.Id, nil

}

// JudgePhone 判断手机号是否存在
func (us *UserService) JudgePhone(phone string) (bool, error) {
	//查询数据库
	var user model.User
	db := tool.GetGormDb()
	var count int64
	if err := db.Model(&model.User{}).Where("phone = ?", phone).First(&user).Count(&count).Error; err != nil {
		if err.Error() == "record not found" {
			return false, nil
		}
		return false, err
	}
	if count == 1 { //找到重复
		return true, nil
	}

	return false, nil

	/*	_,err:=ud.QueryByPhone(phone)
		if err!=nil{
			if err.Error() == "sql: no rows in result set" {
				return false, nil
			}
			return false,err
		}
		return true,nil*/
}

// JudgeVerifyCode 判断验证码是否真确
func (us *UserService) JudgeVerifyCode(ctx *gin.Context, key string, givenValue string) (bool, error) {
	rd := dao.RedisDao{}
	value, err := rd.RedisGetValue(ctx, key)
	if err != nil {

		log.Fatal(err.Error())
		return false, err
	}
	if value != givenValue {
		return false, nil
	}
	return true, nil

}

// InsertUser 将注册成功的用户放入
func (us *UserService) InsertUser(user model.User) error {
	db := tool.GetGormDb()
	if err := db.Create(&user).Error; err != nil {
		log.Fatal(err.Error())
		return err
	}
	/*err:=ud.InsertUser(user)
	if err!=nil{
		return err
	}
	return nil*/
	return nil

}

// LoginByPwd 通过密码登录，返回一个实体
func (us *UserService) LoginByPwd(logins, password string) (model.User, bool, error) {
	var user model.User
	//判断登录类型
	db := tool.GetGormDb()
	index := strings.Index(logins, "@")
	if index != -1 {
		loginname := strings.ToLower(logins)

		if err := db.Model(&model.User{}).Where("email", loginname).First(&user).Error; err != nil {
			if err.Error() == "record not found" {
				return user, false, nil
			}
			return user, false, err
		}
		/*var user, err = ud.QueryByEmail(loginname)
		if err!=nil{
			if err.Error()=="sql: no rows in result set" {
				return model.User{}, false, nil
			}
			return model.User{}, false, err
		}*/
		//使用md5解密
		hash := md5.New()
		hash.Write([]byte(password))
		hash.Write([]byte(user.Salt))
		st := hash.Sum(nil)
		hashPwd := hex.EncodeToString(st)

		if hashPwd != user.Password {
			return model.User{}, false, nil
		}
		return user, true, nil

	} else {
		//手机登录
		if err := db.Model(&model.User{}).Where("phone", logins).First(&user).Error; err != nil {
			if err.Error() == "record not found" {
				return user, false, nil
			}
			return user, false, err
		}
		/*	user,err:=ud.QueryByPhone(logins)
			if err!=nil{
				if err.Error()=="sql: no rows in result set"{
					return model.User{}, false, nil
				}
				return model.User{}, false, err*/
		//使用md5解析
		hash := md5.New()
		hash.Write([]byte(password))
		hash.Write([]byte(user.Salt))
		st := hash.Sum(nil)
		hashPwd := hex.EncodeToString(st)
		if hashPwd != user.Password {
			return model.User{}, false, nil
		}

	}
	return user, true, nil

}
