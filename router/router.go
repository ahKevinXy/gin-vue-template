package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	r = gin.New()

	// the jwt middleware

	// 注册系统路由
	InitSysRouter(r)

	// 注册业务路由
	// TODO: 这里可存放业务路由，里边并无实际路由只有演示代码

	return r
}
