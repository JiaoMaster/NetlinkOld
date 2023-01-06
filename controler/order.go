package controler

import (
	"NetLinkOld/logic"
	"NetLinkOld/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
)

func CreateOrder(c *gin.Context) {
	order := new(models.Order)
	id, err := GetCurrentUser(c)
	err = c.ShouldBindJSON(order)
	cid, _ := strconv.ParseInt(id, 10, 64)
	order.CreateUserId = cid
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
	var re map[string]int
	err := c.ShouldBindJSON(&re)
	ch := re["now"]
	fmt.Println(ch)
	data, err := logic.GetOrderList(page, amount, id, ch)
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

func GetOrderListByOld(c *gin.Context) {
	page := c.Param("page")
	amount := c.Param("amount")
	id, err := GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	data, err := logic.GetOrderListByOld(page, amount, id)
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
	adId := c.Param("adId")
	var payType map[string]int
	err := c.ShouldBindJSON(&payType)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	err = logic.PayOrder(id, payType["payType"], adId)
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
func PayAllOrder(c *gin.Context) {
	var re map[string]interface{}
	err := c.ShouldBindJSON(&re)
	adId := c.Param("adId")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	var oIds []*models.OrderList
	ojson, _ := json.Marshal(re["orders"])
	err = json.Unmarshal(ojson, &oIds)
	payType := re["payType"].(float64)

	var wg sync.WaitGroup
	errCh := make(chan error, len(oIds))
	for _, oId := range oIds {
		id := strconv.FormatInt(oId.Id, 10)
		pay := int(payType)
		wg.Add(1)
		go func(id string, pay int) {
			errCh <- logic.PayOrder(id, pay, adId)
			wg.Done()
		}(id, pay)
	}
	wg.Wait()
	if err := <-errCh; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		close(errCh)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
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
