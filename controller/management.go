package controller

import (
	"net/http"

	"github.com/cargoboat/cargoboat/module/store"
	"github.com/gin-gonic/gin"
)

// configItemModel ...
type configItemModel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Set 添加配置项
func Set(ctx *gin.Context) {
	model := configItemModel{}
	if err := ctx.Bind(&model); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	if err := store.Set(model.Key, model.Value); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.Status(http.StatusOK)
}

// GetAllKeys 获取所有配置项key
func GetAllKeys(ctx *gin.Context) {
	keys := store.GetAllKeys()
	ctx.JSON(http.StatusOK, keys)
}

// GetAll 获取所有配置项
func GetAll(ctx *gin.Context) {
	values := store.GetAll()
	ctx.JSON(http.StatusOK, values)
}
