package handlers

import (
	"net/http"
	"todo/db"

	"github.com/gin-gonic/gin"
)


type User struct {
	Username string  `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"` 
	Email string `json:"email" binding:"required"`
}

func Signup( c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	}
	if err := db.CreateUser(u.Username, u.Password, u.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"user  has been successfully created."})
}

func Update(c *gin.Context) {

	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	}
	if err := db.UpdateUser(u.Username, u.Email, u.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"user successfully has been updated."})
}

