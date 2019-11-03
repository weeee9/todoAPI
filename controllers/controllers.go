package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weeee9/go-todo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type _transTodos struct {
	ID     primitive.ObjectID `json:"_id"`
	Task   string             `json:"task"`
	Status bool               `json:"status"`
}

// NewTask ...
func NewTask(c *gin.Context) {
	var task models.TodoList
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "json binding error:" + err.Error(),
		})
		return
	}
	task.Status = false
	fmt.Println(task)
	if err := models.InsertOneTask(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "task inserted",
	})
}

// GetTasks will get all the tasks
func GetTasks(c *gin.Context) {
	var tasks []*models.TodoList
	var _tasks []_transTodos
	tasks, err := models.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	for _, v := range tasks {
		t := _transTodos{ID: v.ID, Task: v.Task, Status: v.Status}
		_tasks = append(_tasks, t)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"tasks": _tasks,
	})
}

// CompleteTodo will set task status to true
func CompleteTodo(c *gin.Context) {
	_id := c.Param("task")
	if err := models.CompleteTask(_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "task completed",
	})
}

// UndoTodo will set task status to false
func UndoTodo(c *gin.Context) {
	_id := c.Param("task")
	if err := models.UndoTask(_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "task undo",
	})
}

// DeleteOne will delete one task with id
func DeleteOne(c *gin.Context) {
	_id := c.Param("task")
	if err := models.DeleteOneTask(_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "task deleted",
	})
}

// DeleteAll will delete all tasks
func DeleteAll(c *gin.Context) {
	if err := models.DeleteAllTasks(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "all tasks deleted",
	})
}
