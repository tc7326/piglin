package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	//版本1路由
	v1 := r.Group("/v1")
	v1.Use() //token校验中间件

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

}
