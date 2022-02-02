package handlers

import (
	"net/http"
	"todo/todoDB"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	theClaim, err := Auth(c)
	if err != nil {
		return 
	}
	
	title := c.Param("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Please insert the title"})
		return
	}
	

	var todo todoDB.Todo
	todo , err = todoDB.TodoByTitle(theClaim.Username, title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"there is no such a todo"})
		return 
	}
	c.JSON(http.StatusOK, todo)
}