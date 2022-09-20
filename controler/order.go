package controler

import (
	"NetLinkOld/logic"
	"NetLinkOld/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	order := new(models.Order)
	err := c.ShouldBindJSON(order)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	err = logic.CreateOrder(order)
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

func GetOrderList(c *gin.Context) {
	page := c.Param("page")
	amount := c.Param("amount")
	id := c.Param("id")
	data, err := logic.GetOrderList(page, amount, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"msg":       "ok",
		"orderList": data,
	})
}

func GetOrderDetail(c *gin.Context) {
	id := c.Param("id")
	data, err := logic.GetOrderDetail(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "ok",
		"order": data,
	})
}

func PayOrder(c *gin.Context) {
	id := c.Param("id")
	var payType map[string]int
	err := c.ShouldBindJSON(&payType)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	err = logic.PayOrder(id, payType["payType"])
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

func CancelOrder(c *gin.Context) {

}

func UnapplyOrder(c *gin.Context) {
	id := c.Param("id")
	var UnapplyReason map[string]string
	err := c.ShouldBindJSON(&UnapplyReason)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	err = logic.UnapplyOrder(id, UnapplyReason["unapplyReason"])
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
