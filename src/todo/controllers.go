package todo

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var todos []Todo

func GetAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Todos",
		"data":    todos,
	})
	return
}

func CreateTodo(c *gin.Context) {
	var body Todo
	c.BindJSON(&body)
	id := len(todos)
	newTodo := Todo{
		ID:     id,
		Title:  body.Title,
		Body:   body.Body,
		Status: body.Status,
	}

	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, gin.H{
		"id":     newTodo.ID,
		"title":  newTodo.Title,
		"body":   newTodo.Body,
		"status": newTodo.Status,
	})
	return
}

func GetTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("todoId"))

	if err != nil || todoId >= len(todos) {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo not found",
		})
		return
	}

	res := todos[todoId]
	c.JSON(http.StatusOK, gin.H{
		"id":     res.ID,
		"title":  res.Title,
		"body":   res.Body,
		"status": res.Status,
	})
	return
}

func DeleteTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("todoId"))

	if err != nil || todoId >= len(todos) {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo not found",
		})
		return
	}

	todos = append(todos[:todoId], todos[todoId+1:]...)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted successfully",
	})
	return
}
