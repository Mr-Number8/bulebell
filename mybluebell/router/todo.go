package router

import (
	"dist/mybluebell/apis"
	"github.com/gin-gonic/gin"
)

func registerTodoRouter(v1 *gin.RouterGroup,)  {
	api:=apis.TodoApi{}
	r:=v1.Group("/todo") // /v1/todo
	{
		//获取所有待办事项
		r.GET("",api.GetTodoList)
		//获取某个待办事项
		r.GET("/:id",api.GetTodo)
		//
		r.POST("/",api.PostTodo)
		//修改待办事项
		r.PUT("/:id",api.PutTodo)
		//删除
		r.DELETE("/:id", api.DeleteTodo)
	}
}
