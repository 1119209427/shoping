package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shoping/model"
)

// CheckTokenErr 检查并解析token
//token过期返回TOKEN_EXPIRED
//token不正确返回PARSE_TOKEN_ERROR
func CheckTokenErr(ctx *gin.Context, claims *model.MyCustomClaims, err error) bool {
	if err == nil && claims.Type == "ERR" {
		RespErrorWithDate(ctx, "PARSE_TOKEN_ERROR")
		return false
	}

	if err != nil {
		fmt.Println("HE")
		if err.Error()[:16] == "token is expired" {
			RespErrorWithDate(ctx, "TOKEN_EXPIRED")
			return false
		}

		fmt.Println("getTokenParseTokenErr:", err)
		RespErrorWithDate(ctx, "PARSE_TOKEN_ERROR")
		return false
	}

	return true
}
