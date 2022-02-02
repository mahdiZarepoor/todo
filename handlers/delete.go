package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"todo/todoDB"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth(c *gin.Context) (*Claim, error) {
	t, err := c.Cookie("token")

	//checking if there is no token
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "you don't have token , you ARE INVALID"})
			return nil,errors.New("no cookie")
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, errors.New("bar request")
	}

	// checking the token validation and parsing it into variable of type *jwt.Token
	theClaim := &Claim{}
	token, err := jwt.ParseWithClaims(t, theClaim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return nil, errors.New("unauthorized")
		}
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return nil, errors.New("bad request")
	}

	if !token.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the token is invalid"})
		return nil, errors.New("bad request")
	}
	return theClaim,nil 
}

func DeleteTodo(c *gin.Context) {
	theClaim, err := Auth(c)
	if err != nil {
		return 
	}

	title := c.Param("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Please insert the title"})
		return
	}

	r, err := todoDB.DeleteTodo(theClaim.Username, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err":err.Error()})
		return
	}
	if r == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message":"You don't have such a todo"})
		return
	}
	mess := fmt.Sprintf("successfully deleted. %d rows affected", r)
	c.JSON(http.StatusOK, gin.H{"message":mess })

}

func DeleteAllTodos( c *gin.Context) {
	theClaim, err := Auth(c)
	if err != nil {
		return 
	}
	
	rowsAff , err := todoDB.DeleteAll(theClaim.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err":err.Error()})
		return 
	}
	if rowsAff == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Ther is no  todo for this user"})
		return
	}
	mess := fmt.Sprintf("successfully all todos deleted . rows affected %d", rowsAff)
	c.JSON(http.StatusOK, gin.H{"message": mess})
}