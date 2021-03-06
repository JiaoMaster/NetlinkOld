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

	r.POST("/register", controler.RegisterHandler)
	r.POST("/login", controler.LoginHandler)
	apigroup := r.Group("/api")
	//用户路由组
	usergroup := apigroup.Group("/user")
	usergroup.Use(middleware.JWTAuthMiddleware())
	usergroup.POST("/get_user_info", controler.GetUserInfo)
	usergroup.POST("/put_user_info", controler.PutUserInfo)

	//题目路由组

	quegroup := apigroup.Group("/question")
	quegroup.POST("/send_question", middleware.JWTAuthMiddleware(), controler.SendQuestion)
	quegroup.POST("/get_question_detail/:id", controler.GetQuestionDetail)
	quegroup.POST("/get_question_list/:page/:amount", controler.GetQuestionList)

	//评论路由组
	comgroup := apigroup.Group("/commit")
	comgroup.POST("/send_commit", middleware.JWTAuthMiddleware(), controler.SendCommit)
	comgroup.POST("/get_commit/:post_id", controler.GetCommit)

	//版本路由组
	vergroup := apigroup.Group("/version")
	vergroup.POST("/get_version", controler.GetVersion)
	vergroup.POST("/get_apkurl", controler.GetApkUrl)
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
