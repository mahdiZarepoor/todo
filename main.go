package main

import (
	"log"
	"os"
	"todo/handlers"
	"todo/todoDB"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	url    string
	router = gin.Default()
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	url = os.Getenv("SERVER-URL")
}

func main() {

	db := todoDB.GiveMeDb()
	defer db.Close()

	router.POST("/signup",handlers.Signup)
	router.PUT("/update", handlers.Update)
	router.DELETE("/deleteAccount", handlers.DeleteAccount)

	router.POST("/signin", handlers.Signin)
	router.GET("/welcome", handlers.Welcome)
	router.GET("/refresh", handlers.Refresh)

	router.POST("/todo", handlers.CreateTodo)
	router.GET("/todo/:title",handlers.GetTodo)
	router.GET("/todo/list", handlers.ListTodos)
	router.PUT("/todo/done/:title", handlers.Done)  
	router.PUT("todo/undo/:title", handlers.Redo)
	router.DELETE("/todo/:title", handlers.DeleteTodo) 
	router.DELETE("/todo", handlers.DeleteAllTodos) 



	router.Run(url)
}


