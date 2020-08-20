package apis

import (
	"github.com/gin-gonic/gin"

	logs "github.com/sirupsen/logrus"
)

func GetName(c *gin.Context) {
	logs.Error("getName")
	c.JSON(200, "")
}
