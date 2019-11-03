package main

import (
	"github.com/weeee9/go-todo/controllers"

	"github.com/gin-gonic/gin"
	"github.com/weeee9/go-todo/database"
)

func main() {
	router := gin.Default()
	database.Init("./atlas_credentials.json")
	
	router.GET("/todos", controllers.GetTasks)
	router.POST("/todo", controllers.NewTask)
	router.PUT("/todo/:task/complete", controllers.CompleteTodo)
	router.PUT("/todo/:task/undo", controllers.UndoTodo)
	router.DELETE("/todo/:task", controllers.DeleteOne)
	router.DELETE("/todos", controllers.DeleteAll)

	router.Run()
}
