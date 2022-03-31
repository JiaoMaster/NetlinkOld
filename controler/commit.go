package controler

import (
	"NetLinkOld/logic"
	"NetLinkOld/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func SendCommit(c *gin.Context) {
	com := new(models.Commit)
	err := c.ShouldBindJSON(com)
	UserName, err := GetCurrentUser(c)
	com.UserName = UserName
	if err != nil {
		zap.L().Error("controler.SendCommit c.ShouldBindJSON err", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"err":  err,
		})
	}
	err = logic.SendCommit(com)
	if err != nil {
		zap.L().Error("controler.SendCommit logic.SendCommit(com) err", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"err":  err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func GetCommit(c *gin.Context) {
	PostId := c.Param("post_id")
	pId, err := strconv.Atoi(PostId)
	if err != nil {
		zap.L().Error("controler.GetCommit strconv.Atoi err", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"err":  err,
		})
	}
	com, err := logic.GetCommit(pId)
	if err != nil {
		zap.L().Error("controler.GetCommit logic.GetCommit(pId) err", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"err":  err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": com,
	})
}
