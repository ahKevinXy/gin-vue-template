package apis

import (
	"github.com/gin-gonic/gin"


)

func GetName(c *gin.Context) {

	c.JSON(200, "这是我的第一次启动")
}
