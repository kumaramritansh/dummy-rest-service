package routes

import (
	"gitHub.com/apigee/dummy-service/src/todo"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"net/http"
)

func Routes(router *gin.Engine) {
	// get global Monitor object
	m := ginmetrics.GetMonitor()

	m.Use(router)

	// Ping test
	router.GET("/", pingTest)
	router.GET("/ping", pingTest)

	// todo endpoints
	router.GET("/todos", todo.GetAllTodos)
	router.POST("/todo", todo.CreateTodo)
	router.GET("/todo/:todoId", todo.GetTodo)
	router.DELETE("/todo/:todoId", todo.DeleteTodo)
}

func pingTest(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
