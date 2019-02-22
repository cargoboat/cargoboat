package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/cargoboat/cargoboat/module/store"
	"github.com/gin-gonic/gin"
)

// configItemModel ...
type configItemModel struct {
	Key   string `json:"key"`
	Group string `json:"group"`
	Value string `json:"value"`
}

// GetStoreKey 获取存储Key
func (c *configItemModel) GetStoreKey() string {
	return fmt.Sprintf("%s.%s", c.Group, c.Key)
}

// Set 添加配置项
func Set(ctx *gin.Context) {
	model := configItemModel{}
	if err := ctx.Bind(&model); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	//if strings.ToLower(model.Group) == "env" {
	//	ctx.JSON(http.StatusBadRequest, errors.New("group name cannot be env"))
	//}
	if err := store.Set(model.Group, model.Key, model.Value); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.Status(http.StatusOK)
}

// GetAllKeys 获取所有配置项key
func GetAllKeys(ctx *gin.Context) {
	prefix := ctx.Query("prefix")
	prefix = strings.TrimSpace(prefix)
	var keys []string
	if prefix != "" {
		keys = store.GetAllKeysByPrefix(prefix)
	} else {
		keys = store.GetAllKeys()
	}
	ctx.JSON(http.StatusOK, keys)
}

// GetAll 获取所有配置项
func GetAll(ctx *gin.Context) {
	prefix := ctx.Query("prefix")
	prefix = strings.TrimSpace(prefix)
	var values map[string]string
	if prefix != "" {
		values = store.GetAllByPrefix(prefix)
	} else {
		values = store.GetAll()
	}
	ctx.JSON(http.StatusOK, values)
}

// Delete ...
func Delete(ctx *gin.Context) {
	key := ctx.Query("key")
	key = strings.TrimSpace(key)
	if key == "" {
		ctx.JSON(http.StatusBadRequest, errors.New("key cannot be empty"))
		return
	}
	err := store.Delete(key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.Status(http.StatusOK)
}
