package auth

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Authenticator 身份验证
func Authenticator(ctx *gin.Context) (interface{}, error) {
	var loginVals login
	if err := ctx.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	username := loginVals.Username
	password := loginVals.Password
	if username == viper.GetString("system.username") && password == viper.GetString("system.password") {
		return gin.H{
			"user_name": username,
		}, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

// Authorizator 授权
func Authorizator(user interface{}, ctx *gin.Context) bool {
	if v, ok := user.(string); ok && v == viper.GetString("system.username") {
		return true
	}

	return false
}

// Unauthorized 未被授权的
func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
