package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Refresh(c *gin.Context) {
	theClaim, err := Auth(c)
	if err != nil {
		return 
	}

	t, err, expTime := createToken(theClaim.Username, time.Minute*30)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	maxAge := expTime - time.Now().Unix()
	c.SetCookie("token", t, int(maxAge), "/", "localhost", false, true)

}
