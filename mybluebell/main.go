package main

import (
	"dist/mybluebell/models"
	"dist/mybluebell/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func main(){
	//pwd,_:=os.Getwd()
	//fmt.Println("path:",pwd)
	models.InitMsql()
	router.InitRouter()

	
/*
	r:=gin.Default()
	r.Static("/static","static")//引用的静态文件
	r.LoadHTMLGlob("template/*") //加载html模板文件

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	v1Group:=r.Group("/v1")
	{
		//获取所有待办事项
		v1Group.GET("/todo", )
		//获取某个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//添加待办事项
		v1Group.POST("/todo", )
		//修改待办事项
		v1Group.PUT("todo/:id", )
		v1Group.DELETE("todo/:id", func(c *gin.Context) {
			id:=c.Param("id")
			if err:=DB.Where("id=?",id).Error;err!=nil{
				c.JSON(http.StatusOK, gin.H{"error":err.Error()})
			}

			if err:=DB.Where("id=?",id).Delete(&Todo{}).Error;err!=nil{
				c.JSON(http.StatusOK, gin.H{"error":err.Error()})
			}
			c.JSON(http.StatusOK, gin.H{id:"delete"})
		})
	}
	r.Run()
	*/
}