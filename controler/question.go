package controler

import (
	"NetLinkOld/logic"
	"NetLinkOld/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func GetQuestionDetail(c *gin.Context) {
	//绑定参数
	Qid := c.Param("id")
	//业务
	que := new(models.Question)
	var err error
	que, err = logic.GetQuestionDetail(Qid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"code":          202,
		"msg":           "ok",
		"question_info": que,
	})
	return
}

func GetQuestionList(c *gin.Context) {
	//绑定参数
	page := c.Param("page")
	amount := c.Param("amount")
	Page, err := strconv.Atoi(page)
	Amount, err := strconv.Atoi(amount)
	queCh := new(models.QueCh)
	err = c.ShouldBindJSON(queCh)
	if err != nil {
		zap.L().Error(" GetQuestionList 转化失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//逻辑层处理
	data, err := logic.GetQuestionList(Page, Amount, queCh)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//返回响应
	c.JSON(http.StatusOK, gin.H{

		"msg":           "ok",
		"question_list": data,
	})
	return
}

func SendQuestion(c *gin.Context) {
	//绑定
	que := new(models.Question)
	var err error
	err = c.ShouldBindJSON(&que)
	que.UserName, err = GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}

	//业务
	err = logic.SendQuestion(que)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"code": 202,
		"msg":  "ok",
	})
	return
}
