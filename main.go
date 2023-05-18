package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {

	//默认路由
	r := gin.Default()

	//自定义http配置
	s := http.Server{
		Addr:              ":443",
		Handler:           r,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
	}

	//跨域配置
	r.Use(cross)

	//版本1路由
	v1 := r.Group("/v1")
	//token校验中间件
	v1.Use(auth)

	//用户模块
	v1.POST("/register", controller.EventPush) //注册
	v1.POST("/login", controller.EventPush)    //登录
	v1.POST("/logout", controller.EventPush)   //登出

	//相册模块
	v1.GET("/albums", controller.EventPush)       //获取相册列表
	v1.POST("/album", controller.EventPush)       //新建相册
	v1.PUT("/album", controller.EventPush)        //修改相册信息
	v1.DELETE("/album/:id", controller.EventPush) //删除相册

	//图片模块
	v1.GET("/photos", controller.EventPush)       //获取图片列表 条件检索
	v1.POST("/photo", controller.EventPush)       //新增图片
	v1.PUT("/photo", controller.EventPush)        //修改图片信息
	v1.DELETE("/photo/:id", controller.EventPush) //删除图片

	//外链模块

	//启动webserver
	err := s.ListenAndServe()
	if err != nil {
		// Gin 启动失败 终止程序
		panic(err)
	}

}

// cross Gin跨域 https://juejin.cn/post/7032623922902204453
var cross = func(c *gin.Context) {
	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}

// auth token校验
var auth = func(c *gin.Context) {

	url := c.Request.URL.Path
	log.Printf("request url: %s", url)

	//登录接口 不校验token 其余接口都要校验
	if strings.Contains(url, "/manager/login") {
		//直接下一步 不校验
		c.Next()
		return
	}

	//其他接口 校验token
	token := c.GetHeader("Authorization")
	realToken := strings.Split(token, " ")
	if len(realToken) != 2 {
		//token格式异常
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  199,
				"error": "err format of token",
			},
		)
		c.Abort()
		return
	}
	if realToken[0] != "Bearer" {
		//token格式异常
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  199,
				"error": "err format of token",
			},
		)
		c.Abort()
		return
	}
	//token格式正确 从缓存读取token 并比较
	if 1 != 1 {
		//token校验不通过
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  199,
				"error": "err token",
			},
		)
		c.Abort()
		return
	}
	//token 校验通过 放行 进行下一步处理
	c.Next()
	return

}
