package controler

import (
	"NetLinkOld/dao/mysql"
	"NetLinkOld/logic"
	"NetLinkOld/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
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
	username, err := GetCurrentUserName(c)
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

func GetUserNickName(c *gin.Context) {
	//从token获取当前的username
	userid := c.Param("username")
	//username在logic做获取操作
	user, err := mysql.GetUserInfo(userid)
	if err != nil {
		zap.L().Error("logic.GetUserInfo(username) err..", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	//转换结果为json返回
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"user": user,
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
	UserInfo.UserId, _ = strconv.ParseInt(name, 10, 64)
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
	id, err := GetCurrentUser(c)
	err = c.ShouldBindJSON(UserLocation)
	if err != nil {
		zap.L().Error("GetCurrentUser(UserInfo) err..", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	err = logic.PutUserLocation(UserLocation, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
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
			"msg":  err.Error(),
		})
		return
	}
	location, err := logic.GetUserLocation(UserLocation)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
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

func SetUserOld(c *gin.Context) {
	var oldId map[string][]string
	uId, err := GetCurrentUser(c)
	err = c.ShouldBindJSON(&oldId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	oldIds := strings.Join(oldId["oldId"], "+")
	err = mysql.SetOldId(uId, oldIds)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func GetUserOld(c *gin.Context) {
	id, err := GetCurrentUser(c)
	oid, err := mysql.GetOldId(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	oldId := strings.Split(oid, "+")
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "ok",
		"oldId": oldId,
	})
}

func CreateUTS(c *gin.Context) {
	sid := c.Param("shopId")
	uid, ok := c.Get("UserId")
	uidInt := uid.(int64)
	UidS := strconv.FormatInt(uidInt, 10)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "err user",
		})
		return
	}
	err := logic.CreateUserToShop(UidS, sid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func GetUTS(c *gin.Context) {
	id, err := GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	sid, err := logic.GetUTS(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"shopId": sid,
	})
}

func GetUserByOld(c *gin.Context) {
	oid, err := GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	uid, err := logic.GetUserByOld(oid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"userId": uid,
	})
}

func GetOldByUser(c *gin.Context) {
	uid, err := GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	oid, err := logic.GetOldByUser(uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"oldId": oid,
	})
}
