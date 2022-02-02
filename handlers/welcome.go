package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome( c *gin.Context) {
	theClaim, err := Auth(c)
	if err != nil {
		return 
	}
	mess := fmt.Sprintf("Welcome %s", theClaim.Username)
	c.JSON(http.StatusOK, gin.H{"Message":mess})

}

