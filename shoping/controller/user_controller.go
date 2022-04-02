package controller

import (
	"crypto/md5"
	"strings"

	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shoping/model"
	"shoping/param"
	"shoping/service"
	"shoping/tool"
	"strconv"
	"time"
)

type UserController struct {
}

var (
	us = service.UserService{}
	st = service.Tokenservice{}
)

func (uc *UserController) Route(engine *gin.Engine) {
	engine.POST("/api/user/register", uc.register)
	engine.POST("/api/user/login/pw", uc.login)
	engine.POST("/api/verify/email", uc.sendEmailCode)
	engine.GET("/api/check/balance", uc.getBalance)
	engine.GET("/api/check/username", uc.checkUsername)
	engine.GET("/api/check/phone", uc.checkPhone)
	engine.PUT("/api/update/email", uc.updateEmail)
	engine.PUT("/api/update/phone", uc.updatePhone)
	engine.PUT("/api/update/password", uc.updatePwd)
	engine.PUT("/api/update/username", uc.updateUsername)
	engine.PUT("/api/update/gender", uc.updateGender)
	engine.PUT("/api/update/balance", uc.updateBalance)

}

//更新余额
func (uc *UserController) updateBalance(ctx *gin.Context) {
	//获取金额的变化
	change := ctx.PostForm("balance_change")
	if change == "" {
		tool.RespErrorWithDate(ctx, "更改的金额不能为空")
		return

	}

	changeInt, err := strconv.ParseInt(change, 10, 64)
	if err != nil {
		fmt.Println("解析失败")
		tool.RespInternalError(ctx)
		return
	}
	//解析token

	claim, err := ts.ParamToken(ctx.GetHeader("Authorization"))
	if err == nil {
		flag := tool.CheckTokenErr(ctx, claim, err)
		if flag == false {
			return
		}
		//更新用户信息
		user := claim.User
		//判断修改是否合理
		if index := strings.Index(change, "-"); index != -1 {
			//余额减少
			if float64(changeInt) > user.Balance {
				tool.RespErrorWithDate(ctx, "余额不足")
				return

			}

		}

		//修改余额
		err = us.ChangeBalance(changeInt, user.Id)
		if err != nil {
			fmt.Println("修改余额失败")
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessfulWithDate(ctx, "修改余额成功")

	} else {
		tool.RespErrorWithDate(ctx, "token获取失败")
	}

}

//查看余额
func (uc *UserController) getBalance(ctx *gin.Context) {
	//通过id获取余额
	id := ctx.Query("id")
	if id == "" {
		fmt.Println("用户id不能为空")
		tool.RespInternalError(ctx)
		return
	}
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("解析失败")
		tool.RespInternalError(ctx)
		return
	}
	flag, err := us.JudgeUid(uid)
	if err != nil {
		fmt.Println("判断id失败")
		tool.RespInternalError(ctx)
		return
	}
	if flag == false {
		tool.RespErrorWithDate(ctx, "用户id错误")
		return
	}
	balance, err := us.GetBalance(uid)
	if err != nil {
		fmt.Println("获取余额失败")
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx, balance)
	return

}
func (uc *UserController) updateGender(ctx *gin.Context) {
	//1获得token和newgender
	token := ctx.PostForm("token")
	newGender := ctx.PostForm("newGender")
	//解析token
	claim, err := st.ParamToken(token)
	if err != nil {
		fmt.Println("获取参数失败", err)
		tool.RespInternalError(ctx)
		return
	}
	//验证token
	flag := tool.CheckTokenErr(ctx, claim, err)
	if flag == false {
		return
	}
	user := claim.User
	if newGender != "man" && newGender != "women" {
		tool.RespErrorWithDate(ctx, "无效性别")
		return
	}
	//改变性别
	err = us.ChangeGenderById(user.Id, newGender)
	if err != nil {
		fmt.Println("改变性别失败")
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)

}
func (uc *UserController) updateUsername(ctx *gin.Context) {
	//1获取token,newUsername
	token := ctx.PostForm("token")
	newUsername := ctx.PostForm("newUsername")
	//解析token
	claim, err := st.ParamToken(token)
	if err != nil {
		fmt.Println("改变昵称失败", err)
		tool.RespInternalError(ctx)
		return
	}
	user := claim.User
	//验证token
	flag := tool.CheckTokenErr(ctx, claim, err)
	if flag == false {
		tool.RespErrorWithDate(ctx, "token错误")
		return
	}
	if newUsername == "" {
		tool.RespErrorWithDate(ctx, "昵称不能为空")
		return
	}
	if len(newUsername) > 15 {
		tool.RespErrorWithDate(ctx, "昵称太长")
		return
	}
	//改变昵称
	err = us.ChangeUserName(user.Id, newUsername)
	if err != nil {
		fmt.Println("改变昵称失败", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx, "改变昵称成功")

}
func (uc *UserController) updatePwd(ctx *gin.Context) {
	//解析用户表
	var passwords param.ChangePasswordParam
	err := ctx.ShouldBind(&passwords)
	if err != nil {
		tool.RespErrorWithDate(ctx, "解析参数失败")
		return
	}
	//判断账户
	if passwords.Account == "" {
		tool.RespErrorWithDate(ctx, "账户不能为空")
	}
	var id int64
	if index := strings.Index(passwords.Account, "@"); index == -1 {
		//手机号
		flag, err := us.JudgePhone(passwords.Account)
		if err != nil {
			fmt.Println("查询账户失败")
			tool.RespInternalError(ctx)
			return
		}
		if flag == false {
			tool.RespErrorWithDate(ctx, "账户不存在")
			return
		}
		//通过手机号获取id
		id, err = us.SelectIdByPhone(passwords.Account)
		if err != nil {
			fmt.Println("获取id失败")
			tool.RespInternalError(ctx)
			return
		}

	} else {
		//邮箱
		flag, err := us.IsRepeatEmail(passwords.Account)
		if err != nil {
			fmt.Println("查询账户失败")
			tool.RespInternalError(ctx)
			return
		}
		if flag == false {
			tool.RespErrorWithDate(ctx, "账户不存在")
			return
		}
		id, err = us.SelectIdByEmail(passwords.Account)
		if err != nil {
			fmt.Println("获取id失败", err)
			tool.RespInternalError(ctx)
			return
		}
	}
	//验证验证码
	if passwords.Code == "" {
		tool.RespErrorWithDate(ctx, "验证码不能为空")
		return
	}

	flag, err := us.JudgeVerifyCode(ctx, passwords.Account, passwords.Code)
	if err != nil {
		if err.Error() == "redis nil" {
			fmt.Println("未发送验证码")
			tool.RespInternalError(ctx)
			return

		}
		if flag == false {
			tool.RespErrorWithDate(ctx, "验证码输入错误")
			return
		}

	}
	//修改密码
	if passwords.NewPassword == "" {
		tool.RespErrorWithDate(ctx, "密码不能为空")
		return
	}
	if len(passwords.NewPassword) < 7 {
		tool.RespErrorWithDate(ctx, "密码应大于7位")
		return
	}
	err = us.ChangePwd(id, passwords.NewPassword)
	if err != nil {
		fmt.Println("修改密码失败")
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx, "修改密码成功")

}
func (uc *UserController) updatePhone(ctx *gin.Context) {
	var phonepa param.ChangePhoneInfo
	var user model.User
	//解析电话表
	err := ctx.ShouldBind(&phonepa)
	if err != nil {
		fmt.Println("解析失败", err)
		tool.RespInternalError(ctx)
		return
	}
	//解析token
	claim, err := st.ParamToken(phonepa.Token)
	if err != nil {
		fmt.Println("解析token失败", err)
		tool.RespInternalError(ctx)
		return
	}
	//检验token
	flag := tool.CheckTokenErr(ctx, claim, err)
	if flag == false {
		return
	}
	//验证原设备类型
	if index := strings.Index(phonepa.OldAccount, "@"); index == -1 {
		//原设备为手机
		if user.Phone != phonepa.OldAccount {
			tool.RespErrorWithDate(ctx, "输入原设备账号错误，请重新输入")
			return
		}

	} else {
		if user.Email != phonepa.OldAccount {
			tool.RespErrorWithDate(ctx, "原设备账户为空")
			return
		}
		if phonepa.OldCode == "" {
			tool.RespErrorWithDate(ctx, "原设备验证码为空")
			return
		}
	}
	//验证原设备
	flag, err = us.JudgeVerifyCode(ctx, phonepa.OldAccount, phonepa.OldCode)
	if err != nil {
		fmt.Println("获取验证码失败", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag == false {
		tool.RespErrorWithDate(ctx, "验证码错误")
		return
	}
	//验证新设备
	flag, err = us.JudgeVerifyCode(ctx, phonepa.NewPhone, phonepa.NewCode)
	if err != nil {
		fmt.Println("获取验证码失败", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag == false {
		tool.RespErrorWithDate(ctx, "验证码错误")
		return
	}
	//更改手机号
	err = us.ChangePhoneById(user.Id, phonepa.NewPhone)
	if err != nil {
		fmt.Println("修改手机号失败:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)

}
func (uc *UserController) checkPhone(ctx *gin.Context) {
	phone := ctx.Query("phone")
	if phone == "" {
		tool.RespErrorWithDate(ctx, "手机号不能为空，请重新输入")
		return
	}
	flag, err := us.JudgePhone(phone)
	if err != nil {
		fmt.Println("获取手机号失败", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag == true {
		tool.RespErrorWithDate(ctx, "手机号重复，请重新输入")
		return
	}
	tool.RespSuccessful(ctx)
}
func (uc *UserController) updateEmail(ctx *gin.Context) {
	//1解析邮件表
	var emailpa param.ChangeEmailInfo
	err := ctx.ShouldBind(&emailpa)
	if err != nil {
		fmt.Println("参数解析失败:", err)
		tool.RespInternalError(ctx)
		return
	}
	//解析token
	claim, err := st.ParamToken(emailpa.Token)
	//验证token
	flag := tool.CheckTokenErr(ctx, claim, err)
	if flag == false {
		return
	}
	user := claim.User
	//判断原设备
	if index := strings.Index(emailpa.OldAccount, "@"); index == -1 {
		//原设备为手机
		if emailpa.OldAccount != user.Phone {
			tool.RespErrorWithDate(ctx, "原账号错误")
			return
		}

	} else {

		//原设备为email
		if emailpa.OldAccount != user.Email {
			tool.RespErrorWithDate(ctx, "原账号不存在")
			return
		}
		if emailpa.OldCode == "" {
			tool.RespErrorWithDate(ctx, "原账户验证码为空")
			return
		}

	}
	//验证原设备
	flag, err = us.JudgeVerifyCode(ctx, emailpa.OldAccount, emailpa.OldCode)
	if err != nil {
		if err.Error() == "redis: nil" {
			fmt.Println("未发送验证码:", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespInternalError(ctx)
		return
	}
	if flag == false {
		tool.RespErrorWithDate(ctx, "原验证码错误")
		return
	}
	//验证新设备
	flag, err = us.IsRepeatEmail(emailpa.NewEmail)
	if err != nil {
		fmt.Println("获取邮箱失败", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag == true {
		tool.RespErrorWithDate(ctx, "邮箱重复")
		return
	}
	flag, err = us.JudgeVerifyCode(ctx, emailpa.NewEmail, emailpa.NewCode)
	if err != nil {
		if err.Error() == "redis nil" {
			fmt.Println("未发生验证码:", err)
			tool.RespInternalError(ctx)
			return
		}
		fmt.Println("获取验证码失败", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag == true {
		tool.RespErrorWithDate(ctx, "验证码错误")
		return
	}
	//修改邮箱
	err = us.ChangeEmail(user.Id, emailpa.NewEmail)
	if err != nil {
		fmt.Println("修改邮箱失败:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)

}
func (uc *UserController) sendEmailCode(ctx *gin.Context) {
	email := ctx.PostForm("email")
	code, err := us.SendEmailCode(email)
	if err != nil {
		if err.Error()[:3] == "cod" {
			tool.RespErrorWithDate(ctx, "邮箱不正确")
			return
		}
		fmt.Println("发送邮件失败")
		tool.RespInternalError(ctx)
		return
	}
	//将验证码存入redis
	err = us.VerifyCodeIn(ctx, email, code)
	if err != nil {
		fmt.Println("发送邮件出错:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)

}

//登录
func (uc *UserController) login(ctx *gin.Context) {
	logins := ctx.PostForm("logins")
	password := ctx.PostForm("password")
	if logins == "" {
		tool.RespErrorWithDate(ctx, "用户名不能为空")
		return
	}
	if password == "" {
		tool.RespErrorWithDate(ctx, "密码不能为空")
		return
	}
	user, flag, err := us.LoginByPwd(logins, password)
	if err != nil {
		fmt.Println("登录服务出错")
		tool.RespInternalError(ctx)
		return
	}
	if flag == false {
		tool.RespErrorWithDate(ctx, "密码错误，请重新输入")
		return
	}
	//创建token 2分钟
	token, err := st.CreateToken(user, 120, "TOKEN")
	if err != nil {
		fmt.Println("创建token失败")
		tool.RespInternalError(ctx)
		return
	}
	//创建refreshtoken 1天
	retoken, err := st.CreateToken(user, 86400, "REFRESHTOKEM")
	if err != nil {
		fmt.Println("创建refreshtoken失败")
		tool.RespInternalError(ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":       true,
		"data":         user.Id,
		"token":        token,
		"refreshToken": retoken,
		"uid":          user.Id,
	})

	ctx.SetCookie("loginname", logins, 600, "/", "", false, false)

}

//检查用户名是否符合要求
func (uc *UserController) checkUsername(ctx *gin.Context) {
	username := ctx.Query("username")
	if username == "" {
		tool.RespErrorWithDate(ctx, "请告诉我你的昵称吧")
		return
	}

	if len(username) > 14 {
		tool.RespErrorWithDate(ctx, "昵称过长")
		return
	}
	flag, err := us.IsRepeatName(username)
	if err != nil {
		fmt.Println("查询出错")
		tool.RespInternalError(ctx)
		return
	}
	if flag {
		tool.RespErrorWithDate(ctx, "昵称重复，请重新输入")
		return
	}
	tool.RespSuccessful(ctx)

}
func (uc *UserController) register(ctx *gin.Context) {
	//1.解析用户表
	var userparam = param.UserParam{}
	err := ctx.ShouldBind(&userparam)
	if err != nil {
		tool.RespErrorWithDate(ctx, "参数解析失败")
		return
	}
	//2.判断用户名是否符合要求
	if len(userparam.Username) > 10 {
		tool.RespErrorWithDate(ctx, "用户名过长，不能大于10位")
		return
	}
	if len(userparam.Username) == 0 {
		tool.RespErrorWithDate(ctx, "用户名不能为空，请重新输入")
		return
	}

	//3.判断密码
	if len(userparam.Pwd) < 7 {
		tool.RespErrorWithDate(ctx, "密码太短，长度必须大于7位")
		return
	}
	if len(userparam.Pwd) == 0 {
		tool.RespErrorWithDate(ctx, "密码不能为空，请重新输入")
		return
	}
	if len(userparam.Pwd) > 20 {
		tool.RespErrorWithDate(ctx, "密码太长了，不能大于20位")
		return
	}
	//判断手机号
	phone := userparam.Phone

	flag, err := us.JudgePhone(phone)

	if err != nil {
		fmt.Println("判断手机号出错:", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag == true {
		tool.RespErrorWithDate(ctx, "手机号已经存在，请重新输入")
		return

	}
	//判断验证码
	/*givenCode := userparam.VerifyCode
	flag,err=us.JudgeVerifyCode(ctx,phone,givenCode)
	if err!=nil{
		if err.Error() == "redis: nil"{
			fmt.Println("未发送验证码")
			tool.RespInternalError(ctx)
			return
		}
		fmt.Println("判断验证码出错",err)
		tool.RespInternalError(ctx)
		return
	}
	if flag==false{
		tool.RespErrorWithDate(ctx,"验证码错误，请重新输入")
		return
	}*/
	var user = model.User{}
	//user.RegDate=time.Now() gorm.model 有这个属性
	user.Phone = phone
	user.Username = userparam.Username
	//加密
	user.Salt = strconv.FormatInt(time.Now().Unix(), 10)
	m5 := md5.New()
	m5.Write([]byte(userparam.Pwd))
	m5.Write([]byte(user.Salt))
	st := m5.Sum(nil)
	user.Password = hex.EncodeToString(st)
	//放入mysql数据库
	err = us.InsertUser(user)
	if err != nil {
		fmt.Println("数据存储失败")
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}
