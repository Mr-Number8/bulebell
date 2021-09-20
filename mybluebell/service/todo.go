package service

import (
	"dist/mybluebell/models"
)

/*
*/
type TodoService struct {

}
//获取所有待办事项
func (s TodoService)GetTodoList()([]models.Todo,error) {
	var todoList []models.Todo
	err:=models.DB.Find(&todoList).Error
	return todoList,err
}
//获取一个待办事项
func (s TodoService)GetTodo(id string)(models.Todo,error) {
	var todo  models.Todo
	err:=models.DB.Where("id=?",id).Find(&todo).Error
	return todo,err
}
//添加待办事项
func (s TodoService)PostTodo(todo models.Todo)error  {
	return models.DB.Create(&todo).Error;
}
//修改待办事项
func (s TodoService)PutTodo(todo models.Todo)(error ){
	//var todo models.Todo
	//查询数据库是否有这个事项
	//id,_:=c.Params.Get("id")
	/*if err:=models.DB.Where("id=?",id).First(&todo).Error;err!=nil{
		//c.JSON(http.StatusOK, gin.H{"error":err.Error()})
		return err,nil
	}*/
	//获取修改后的数据
	//c.BindJSON(&todo)
	//保存到数据库
	//todo.Status= status
	if err:=models.DB.Save(&todo).Error;err!=nil{
		//c.JSON(http.StatusOK, gin.H{"error":err.Error()})
		return err
	}
	return nil
	//c.JSON(http.StatusOK, todo)
}

func (s TodoService)DeleteTodo(id string)error{
	/*id:=c.Param("id")
	if err:=DB.Where("id=?",id).Error;err!=nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}*/
	//
	_,err:=s.GetTodo(id)
	if err!=nil{
		return err
	}
	return models.DB.Where("id=?",id).Delete(&models.Todo{}).Error

	//c.JSON(http.StatusOK, gin.H{id:"delete"})
}
