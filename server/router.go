package server

import (
	"net/http"

	"github.com/cargoboat/cargoboat/controller"

	"github.com/cargoboat/cargoboat/controller/client"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func getClientAccounts() gin.Accounts {
	username := viper.GetString("client.basic_auth.username")
	password := viper.GetString("client.basic_auth.password")
	return gin.Accounts{
		username: password,
	}
}
func getServerAccounts() gin.Accounts {
	username := viper.GetString("server.basic_auth.username")
	password := viper.GetString("server.basic_auth.password")
	return gin.Accounts{
		username: password,
	}
}

// setRouter 设置路由
func setRouter(handler *gin.Engine) {
	handler.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "welcome cargoboat server")
	})

	serverAuth := handler.Group("/", gin.BasicAuth(getServerAccounts()))
	{
		serverAuth.POST("/set", controller.Set)
		serverAuth.GET("/keys", controller.GetAllKeys)
		serverAuth.GET("/all", controller.GetAll)
	}

	clientAuth := handler.Group("/client", gin.BasicAuth(getClientAccounts()))
	{
		clientAuth.GET("/version", client.GetVersion)
		clientAuth.GET("/configs", client.Get)
	}
}
