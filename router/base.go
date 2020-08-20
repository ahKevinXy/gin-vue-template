package router

import (
	"github.com/gin-gonic/gin"
	"learn-go/api/app"
	"mime"
)

func InitSysRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("")
	sysBaseRouter(g)
	// 静态文件
	sysStaticFileRouter(g)

	// 无需认证
	sysNoCheckRoleRouter(g)
	// 需要认证
	sysCheckRoleRouterInit(g)

	return g
}

func sysBaseRouter(r *gin.RouterGroup) {

}

func sysStaticFileRouter(r *gin.RouterGroup) {
	mime.AddExtensionType(".js", "application/javascript")

	r.Static("/static", "./static")
	r.Static("/form-generator", "./static/form-generator")
}

func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	v1 := r.Group("/")
	v1.GET("/", app.GetName)
	registerPublicRouter(v1)

}

func sysCheckRoleRouterInit(r *gin.RouterGroup) {

}

func registerPublicRouter(v1 *gin.RouterGroup) {

}
