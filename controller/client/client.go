package client

import (
	"net/http"

	"github.com/cargoboat/cargoboat/module/store"
	"github.com/gin-gonic/gin"
)

// configItem ...
type configItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Get 获取
func Get(ctx *gin.Context) {
	items := make([]configItem, 0)
	groupName := ctx.MustGet(gin.AuthUserKey).(string)
	for key, value := range store.GetAllByPrefix(groupName) {
		items = append(items, configItem{
			Key:   key,
			Value: value,
		})
	}
	for key, value := range store.GetAllByPrefix("env") {
		items = append(items, configItem{
			Key:   key,
			Value: value,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"version": store.GetVersion(),
		"configs": items,
	})
}

// GetVersion 获取版本
func GetVersion(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"version": store.GetVersion(),
	})
}
