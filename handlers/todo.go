package handlers

import (
	"fmt"
	"net/http"
	"todo/todoDB"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Deadline    string `json:"deadline" binding:"required"`
}

func CreateTodo(c *gin.Context) {
	theClaim, err := Auth(c)
	if err != nil {
		return 
	}


	var t Todo
	if err := c.ShouldBindJSON(&t); err != nil {
		mess := fmt.Sprintf("can't bind the given data to our model : %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": mess})
	}
	

	err =todoDB.InsertTodo(theClaim.Username,t.Title, t.Description, t.Deadline)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"This todo already exists for this user!"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message":"successfully added"})
}

