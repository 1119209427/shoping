package service

import (
	"github.com/dgrijalva/jwt-go"
	"shoping/model"
	"shoping/tool"
	"time"
)

type Tokenservice struct{

}

// ParamRefreshToken 解析refreshtoken
func(ts *Tokenservice)ParamRefreshToken(tokenstring string) (*model.MyCustomClaims, error){
	jwtcfg:=tool.GetCfg().Jwt
	mySigningKey:=[]byte(jwtcfg.SigningKey)
	token,err:=jwt.ParseWithClaims(tokenstring, &model.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if clams, ok := token.Claims.(*model.MyCustomClaims); ok && token.Valid {
		if clams.Type == "REFRESH_TOKEN" {
			errClaims := new(model.MyCustomClaims)
			errClaims.Type = "ERR"
			return errClaims, nil
		}
		return clams, nil
	} else {
		return nil, err
	}

}

// ParamToken 解析token
func(ts *Tokenservice)ParamToken(tokenstring string) (*model.MyCustomClaims, error) {
	jwtcfg:=tool.GetCfg().Jwt
	mySigningKey:=[]byte(jwtcfg.SigningKey)
	 token,err:=jwt.ParseWithClaims(tokenstring, &model.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if clams, ok := token.Claims.(*model.MyCustomClaims); ok && token.Valid {
		if clams.Type == "TOKEN" {
			errClaims := new(model.MyCustomClaims)
			errClaims.Type = "ERR"
			return errClaims, nil
		}
		return clams, nil
	} else {
		return nil, err
	}

}

// CreateToken 创建jwt
func(ts *Tokenservice)CreateToken(user model.User,ExpireTime int64, tokenType string)(string,error){
jwtcfg:=tool.GetCfg().Jwt
mySigningKey:=[]byte(jwtcfg.SigningKey)
claims:=model.MyCustomClaims{
	User: user,
	Type: tokenType,
StandardClaims:jwt.StandardClaims{
	ExpiresAt: time.Now().Unix() + ExpireTime,
},
}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)


}
