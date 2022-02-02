package handlers

import (
	"net/http"
	"todo/todoDB"

	"github.com/gin-gonic/gin"
)

func ListTodos(c *gin.Context) {
	theClaim, err := Auth(c)
	if err != nil {
		return 
	}
	
	var todos []todoDB.Todo 
	todos, err = todoDB.ListAllTodos(theClaim.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	if todos == nil {
		c.JSON(http.StatusOK, gin.H{"message" : "empty todo"})
		return
	}
	c.JSON(http.StatusOK, todos)
}