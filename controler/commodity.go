package controler

import (
	"NetLinkOld/dao/mysql"
	"NetLinkOld/logic"
	"NetLinkOld/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTypeList(c *gin.Context) {
	sqlStr := "select id, name, image from commodityType"
	tL := []*models.CommodityType{}
	mysql.Db.Select(&tL, sqlStr)
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"TypeList": tL,
	})
}

func CreateCommodity(c *gin.Context) {
	com := new(models.Commodity)
	err := c.ShouldBindJSON(com)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	err = logic.CreateCommodity(com)
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

func GetCommodityList(c *gin.Context) {
	page := c.Param("page")
	amount := c.Param("amount")
	typeid := c.Param("type")
	r, err := logic.GetCommodityList(page, amount, typeid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":          200,
		"msg":           "ok",
		"commodityList": r,
	})
}

func GetCommodityDetail(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.ParseInt(id, 10, 64)
	re, err := logic.GetCommodityDetail(idInt)
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
		"shop": re,
	})
}
