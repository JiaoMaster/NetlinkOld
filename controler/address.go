package controler

import (
	"NetLinkOld/logic"
	"NetLinkOld/models"
	"github.com/gin-gonic/gin"
)

func AddAd(c *gin.Context) {
	uid := c.Param("uid")
	ad := new(models.Address)
	err := c.ShouldBindJSON(ad)
	if err != nil {
		c.JSON(200, gin.H{
			"code": "404",
			"err":  err.Error(),
		})
	}
	err = logic.AddAddress(uid, *ad)
	if err != nil {
		c.JSON(200, gin.H{
			"code": "404",
			"err":  err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code": "200",
	})

}

func DelAd(c *gin.Context) {
	id := c.Param("id")
	err := logic.DelAddress(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": "404",
			"err":  err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code": "200",
	})
}

func UpDateAd(c *gin.Context) {
	id := c.Param("id")
	ad := new(models.Address)
	err := c.ShouldBindJSON(ad)
	if err != nil {
		c.JSON(200, gin.H{
			"code": "404",
			"err":  err.Error(),
		})
	}
	err = logic.UpdateAddress(id, *ad)
	if err != nil {
		c.JSON(200, gin.H{
			"code": "404",
			"err":  err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code": "200",
	})

}

func GetAd(c *gin.Context) {
	id := c.Param("id")
	data, err := logic.GetAd(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": "404",
			"err":  err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"address": data,
	})
}

func GetAdList(c *gin.Context) {
	uid := c.Param("uid")
	data, err := logic.GetAdList(uid)
	if err != nil {
		c.JSON(200, gin.H{
			"code": "404",
			"err":  err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":        "200",
		"addressList": data,
	})
}
