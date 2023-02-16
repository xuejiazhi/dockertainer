package main

import (
	"dockertainer/api/cache"
	"dockertainer/api/endpoint"
	"dockertainer/api/user"
	"dockertainer/api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	MacAddr string
)

func Route() {
	MacAddr = util.GetMac()
	r := gin.Default() //创建一个默认的路由引擎

	//登录校验
	login := r.Group("/login")
	{
		login.POST("/auth", user.Auth)
	}

	//用户操作
	users := r.Group("/user")
	users.Use(MiddleWare)
	{
		users.POST("/modifypass", user.ModifyPassWord)
	}

	//node 操作
	node := r.Group("/node")
	node.Use(MiddleWare)
	{
		node.POST("/list", endpoint.NodeList)
		node.POST("/add", endpoint.AddNode)
		node.GET("/modify", endpoint.ModifyNode)
		node.GET("/remove", endpoint.RemoveNode)
	}

	//面版
	dashboard := r.Group("/dashboard")
	dashboard.Use(MiddleWare)
	{
		dashboard.GET("/info", endpoint.GetDashBoardInfo)
	}

	//镜像操作
	images := r.Group("/images")
	images.Use(MiddleWare)
	{
		images.GET("/list", endpoint.ImagesJson)
		images.GET("/remove", endpoint.RemoveImage)
		images.GET("/inspect", endpoint.InspectImage)
	}

	r.Run()
}

func MiddleWare(c *gin.Context) {
	if _, ok := cache.Get(MacAddr); ok {
		c.Next()
	} else {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
		// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
		return
	}
}
