package apis

import (
	"dist/mybluebell/models"
	"dist/mybluebell/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoApi struct {

}

func (a *TodoApi)GetTodoList(c *gin.Context){
	s:=service.TodoService{}
	todoList,err:=s.GetTodoList()
	if err!=nil{
	  	c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	  }else{
	  	c.JSON(http.StatusOK,todoList)
	  }
}
func (a *TodoApi)GetTodo(c *gin.Context){
	s:=service.TodoService{}
	id:=c.Param("id")
	todo,err:=s.GetTodo(id)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,todo)
	}
}

func(a *TodoApi)PostTodo(c *gin.Context) {
	s:=service.TodoService{}
	var todo models.Todo
	c.BindJSON(&todo)
	err:=s.PostTodo(todo)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK, todo)
	}

}
func(a *TodoApi)PutTodo(c *gin.Context){
	s:=service.TodoService{}
	var todo models.Todo

	id:=c.Param("id")
	//var status bool
	todo,err:=s.GetTodo(id)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}
	c.BindJSON(&todo)
	if err:=s.PutTodo(todo);err!=nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK, todo)
	}
}

func(a *TodoApi)DeleteTodo(c *gin.Context){
	s:=service.TodoService{}
	//var todo models.Todo
	id:=c.Param("id")
	if err:=s.DeleteTodo(id);err!=nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{id:"delete"})
	}
}