package server

import (
	"net/http"

	"github.com/spf13/viper"

	"github.com/cargoboat/cargoboat/controller/config"
	"github.com/gin-gonic/gin"
)

func getAccounts() gin.Accounts {
	username := viper.GetString("basic_auth.username")
	password := viper.GetString("basic_auth.password")
	return gin.Accounts{
		username: password,
	}
}

// setRouter 设置路由
func setRouter(handler *gin.Engine) {
	handler.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "welcome cargoboat server")
	})

	auth := handler.Group("/client", gin.BasicAuth(getAccounts()))
	{
		auth.GET("/version", config.GetVersion)
		auth.GET("/configs", config.Get)
	}
}
