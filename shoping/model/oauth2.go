
package model
type Oauth2 struct {
	ClientId uint
	RedirectUrl string
	Login string//提供用于登录和授权应用程序的特定账户
	Scopes []string
	State string
}
