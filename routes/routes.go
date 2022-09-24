package routes

import (
	"NetLinkOld/controler"
	"NetLinkOld/logger"
	"NetLinkOld/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(Cors())
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	})

	//注册登录控制组
	{
		r.POST("/register", controler.RegisterHandler)
		r.POST("/login", controler.LoginHandler)
	}
	apigroup := r.Group("/api")
	{
		//用户路由组
		usergroup := apigroup.Group("/user")
		{
			usergroup.POST("/get_user_NickName/:username", controler.GetUserNickName)
			usergroup.Use(middleware.JWTAuthMiddleware())
			usergroup.POST("/get_user_info", controler.GetUserInfo)
			usergroup.POST("/put_user_info", controler.PutUserInfo)
			usergroup.POST("/get_user_name/:id", controler.GetUserName)
			usergroup.POST("/get_user_location", controler.GetUserLocation)
			usergroup.POST("/put_user_location", controler.PutUserLocation)
			usergroup.POST("/set_old/:id", controler.SetUserOld)
			usergroup.POST("/get_old/:id", controler.GetUserOld)
		}
		//问题路由组

		quegroup := apigroup.Group("/question")
		{
			quegroup.POST("/send_question", middleware.JWTAuthMiddleware(), controler.SendQuestion)
			quegroup.POST("/get_question_detail/:id", controler.GetQuestionDetail)
			quegroup.POST("/get_question_list/:page/:amount", controler.GetQuestionList)
			quegroup.POST("/get_question_by_id/:id/:page/:amount", controler.GetQuestionListById)
		}

		//评论路由组
		comgroup := apigroup.Group("/commit")
		{
			comgroup.POST("/send_commit", middleware.JWTAuthMiddleware(), controler.SendCommit)
			comgroup.POST("/get_commit/:post_id", controler.GetCommit)
		}

		//版本路由组
		vergroup := apigroup.Group("/version")
		{
			vergroup.POST("/get_version", controler.GetVersion)
			vergroup.POST("/get_apkurl", controler.GetApkUrl)
		}

		//通知websocket组
		nroup := apigroup.Group("/notice")
		{
			nroup.GET("/newcommit", controler.NewCommit)
		}

		//商铺组
		shop := apigroup.Group("/shop")
		{
			shop.POST("/create", controler.CreateShop)
			shop.POST("/GetList/:page/:amount/:type", controler.GetShopList)
			shop.POST("/GetDetail/:id", controler.GetShopDetail)
		}

		//商品组
		commodity := apigroup.Group("/commodity")
		{
			commodity.POST("/create", controler.CreateCommodity)
			commodity.POST("/GetTypeList", controler.GetTypeList)
			commodity.POST("/GetList/:page/:amount/:type", controler.GetCommodityList)
			commodity.POST("/GetDetail/:id", controler.GetCommodityDetail)
		}

		//订单组
		order := apigroup.Group("/order")
		{
			order.Use(middleware.JWTAuthMiddleware())
			order.POST("/create", controler.CreateOrder)
			order.POST("/GetList/:page/:amount/:id", controler.GetOrderList)
			order.POST("/GetDetail/:id", controler.GetOrderDetail)
			order.POST("/Pay/:id", controler.PayOrder)
			order.POST("/cancel/:id", controler.CancelOrder)
			order.POST("/unapply/:id", controler.UnapplyOrder)
		}
		//address
		ad := apigroup.Group("/address")
		{
			ad.Use(middleware.JWTAuthMiddleware())
			ad.POST("/add/:uid", controler.AddAd)
			ad.POST("/del/:id", controler.DelAd)
			ad.POST("/update/:id", controler.UpDateAd)
			ad.POST("/get/:id", controler.GetAd)
			ad.POST("/getList/:uid", controler.GetAdList)
		}
	}

	r.POST("/upload/:ch/:location", middleware.JWTAuthMiddleware(), controler.SendAudioQue)
	r.GET("/down", controler.DownloadFileControl)

	// 文件上传组
	{
		r.POST("/objects/:filename", func(context *gin.Context) {
			controler.Handler(context.Writer, context.Request)
		})
		r.GET("/objects/:filename", func(context *gin.Context) {
			controler.Handler(context.Writer, context.Request)
		})
		r.DELETE("/objects/:filename", func(context *gin.Context) {
			controler.Handler(context.Writer, context.Request)
		})
	}
	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
