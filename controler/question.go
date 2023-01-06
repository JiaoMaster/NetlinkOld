package controler

import (
	"NetLinkOld/logic"
	"NetLinkOld/models"
	"NetLinkOld/pkg/uuid"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"net/http"
	"os"
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
	que.UserName, err = GetCurrentUserName(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
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

func GetQuestionListById(c *gin.Context) {
	//绑定参数
	id := c.Param("id")
	page := c.Param("page")
	amount := c.Param("amount")
	Page, err := strconv.Atoi(page)
	Amount, err := strconv.Atoi(amount)
	if err != nil {
		zap.L().Error(" GetQuestionList 转化失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	//逻辑层处理
	data, err := logic.GetQuestionListById(id, Page, Amount)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
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

func SendAudioQue(c *gin.Context) {
	que := new(models.Question)
	var err error
	ch := c.Param("ch")
	location := c.Param("location")
	que.CommunityID, _ = strconv.ParseInt(ch, 10, 64)
	que.Location = location
	que.UserName, err = GetCurrentUserName(c)

	// 单个文件
	file, err := c.FormFile("picFile")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	ID, err := uuid.Getuuid()
	// 上传文件到指定的目录
	dst := fmt.Sprintf("%s/%s", viper.GetString("tmp.audio_path"), strconv.FormatInt(ID, 10)+".mp3")
	err = c.SaveUploadedFile(file, dst)

	que.ID = ID
	que.AudioPath = dst
	err = logic.SendAudioQue(que)
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
}

func DownloadFileControl(c *gin.Context) {

	type Path struct {
		path string `json:"path"`
	}
	// 获取要返回的文件数据流
	var ppath = Path{}
	ppath.path = c.Query("url")

	mp3, err := os.OpenFile(ppath.path, os.O_RDWR, os.ModeTemporary)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	defer mp3.Close()

	// 设置返回头并返回数据
	c.Header("Content-Type", "audio/mpeg")
	c.File(ppath.path)
	c.JSON(http.StatusOK, gin.H{
		"code": 202,
		"msg":  "ok",
	})
}
