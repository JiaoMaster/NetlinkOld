package controler

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strconv"
)

func GetCurrentUser(c *gin.Context) (string, error) {
	data, ok := c.Get("UserId")
	if !ok {
		zap.L().Error("GetCurrentUser err...", zap.Error(errors.New("用户获取失败")))
		return "", errors.New("用户获取失败")
	}
	user := data.(int64)
	uid := strconv.FormatInt(user, 10)
	return uid, nil
}
