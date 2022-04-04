package controler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func GetVersion(c *gin.Context) {
	version := fmt.Sprintf("%s",
		viper.GetString("app.version"),
	)

	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"code": 202,
		"msg":  "ok",
		"data": version,
	})
	return
}

func GetApkUrl(c *gin.Context) {
	url := fmt.Sprintf("%s",
		viper.GetString("app.apkurl"),
	)

	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"code": 202,
		"msg":  "ok",
		"data": url,
	})
	return
}
