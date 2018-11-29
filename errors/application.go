package errors

import "errors"

// Application
var (
	// ErrApplicationNotFound 未查到配置项
	ErrApplicationNotFound = errors.New("未查到配置项")
	// ErrApplicationCreateFailure 创建应用失败
	ErrApplicationCreateFailure = errors.New("创建应用失败")
	// ErrApplicationDelFailure 删除应用失败
	ErrApplicationDelFailure = errors.New("删除应用失败")
)
