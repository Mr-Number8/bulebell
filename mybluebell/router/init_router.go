package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(){
	r:=gin.Default()
	r.Static("/static","static")//引用的静态文件
	//E:\go-project\go-bluebell\dist\mybluebell\template
	r.LoadHTMLGlob("dist/mybluebell/template/*") //加载html模板文件
	v1Group:=r.Group("v1")
	//注册todo
	registerTodoRouter(v1Group)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	r.Run()
}
