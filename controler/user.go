package controler

import (
	"NetLinkOld/dao/mysql"
	"NetLinkOld/logic"
	"NetLinkOld/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func RegisterHandler(c *gin.Context) {
	user := new(models.UserSignUp)
	if err := c.ShouldBindJSON(user); err != nil {
		zap.L().Error("ShouldBindJSON(user) err...", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//在logic层实现注册
	err := logic.Register(user)
	if err != nil {
		zap.L().Error("logic.Register(user) err...", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func LoginHandler(c *gin.Context) {
	//绑定参数
	data := new(models.UserSignUp)
	err := c.ShouldBindJSON(data)
	if err != nil {
		zap.L().Error("LoginHandler ShouldBindJSON(data) err...", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//交给logic走逻辑业务
	token, err := logic.Login(data)
	if err != nil {
		zap.L().Error("logic.Login(data) err...", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "ok",
		"token": token,
	})
	return
}

func GetUserInfo(c *gin.Context) {
	//从token获取当前的username
	username, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("GetCurrentUser(c) err..", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//username在logic做获取操作
	user, err := logic.GetUserInfo(username)
	if err != nil {
		zap.L().Error("logic.GetUserInfo(username) err..", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//转换结果为json返回
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"info": user,
	})
	return
}

func GetUserName(c *gin.Context) {
	//从token获取当前的username
	userid := c.Param("id")
	id, err := strconv.ParseInt(userid, 10, 64)
	//username在logic做获取操作
	username, err := mysql.GetUsername(id)
	if err != nil {
		zap.L().Error("logic.GetUserInfo(username) err..", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//转换结果为json返回
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"name": username,
	})
	return
}

func PutUserInfo(c *gin.Context) {
	UserInfo := new(models.User)
	err := c.ShouldBindJSON(UserInfo)
	if err != nil {
		zap.L().Error("GetCurrentUser(UserInfo) err..", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	name, err := GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	UserInfo.Username = name
	err = logic.PutUserInfo(UserInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
	return
}

func PutUserLocation(c *gin.Context) {
	UserLocation := new(models.UserLocation)
	err := c.ShouldBindJSON(UserLocation)
	if err != nil {
		zap.L().Error("GetCurrentUser(UserInfo) err..", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	err = logic.PutUserLocation(UserLocation)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
	return
}

func GetUserLocation(c *gin.Context) {
	UserLocation := new(models.UserLocation)
	err := c.ShouldBindJSON(UserLocation)
	if err != nil {
		zap.L().Error("GetCurrentUser(UserInfo) err..", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	location, err := logic.GetUserLocation(UserLocation)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "ok",
		"location": location,
	})
	return
}
