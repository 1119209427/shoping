package param
type UserParam struct {
	Id   int64`form:"id"`
	Username   string `form:"username"`
	Pwd        string `form:"password"`
	Phone      string `form:"phone"`
	VerifyCode string `form:"verify_code"`
}
