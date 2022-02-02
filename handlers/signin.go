package handlers

import (
	"log"
	"net/http"
	"os"
	"time"
	"todo/db"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	jwtKey []byte
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	jwtKey = []byte(os.Getenv("JWT_KEY"))

}

func Signin(c *gin.Context) {
	var u Credential
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	valid, err := db.IsValid(u.Username, u.Password)
	if valid == false {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User invalid"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err, expTime := createToken(u.Username, time.Second*30)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	maxAge := expTime - time.Now().Unix()
	c.SetCookie("token", token, int(maxAge), "/", "localhost", false, true)

}

func createToken(user string, d time.Duration) (string, error, int64) {
	expTime := time.Now().Add(d).Unix()
	claim := &Claim{
		Username: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := t.SignedString(jwtKey)
	if err != nil {
		return "", err, 0
	}
	return token, nil, expTime
}

func DeleteAccount(c *gin.Context) {
	theClaim, err := Auth(c)
	if err != nil {
		return 
	}


	if err := db.DeleteAccount(theClaim.Username); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"user has been successfully deleted"})

}
