package handlers

import (
	"net/http"
	"todo/todoDB"

	"github.com/gin-gonic/gin"
)

func Redo(c *gin.Context) {
	theClaim, err := Auth(c)
	if err != nil {
		return 
	}

	title := c.Param("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Please insert the title"})
		return
	}

	rowsAff, err := todoDB.UndoTodo(theClaim.Username , title )
	if err != nil  {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	if rowsAff == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message":"You have no such a todo"})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message":"Undid"})
}


