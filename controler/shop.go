package controler

import (
	"NetLinkOld/dao/mysql"
	"NetLinkOld/logic"
	"NetLinkOld/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetShopTypeList(c *gin.Context) {
	sqlStr := "select id, name, image from shop"
	tL := []*models.ShopType{}
	mysql.Db.Select(&tL, sqlStr)
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"TypeList": tL,
	})
}

func CreateShop(c *gin.Context) {
	shop := new(models.Shop)
	err := c.ShouldBindJSON(shop)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	id, err := logic.CreateShop(shop)
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
		"data": id,
	})
}

func GetShopList(c *gin.Context) {
	page := c.Param("page")
	amount := c.Param("amount")
	typeid := c.Param("type")
	r, err := logic.GetShopList(page, amount, typeid)
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
		"shopList": r,
	})
}

func GetShopDetail(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.ParseInt(id, 10, 64)
	re, err := logic.GetShopDetail(idInt)
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
