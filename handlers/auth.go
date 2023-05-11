package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"go-practice/db"
)

func loginHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	fmt.Println("email: ", email)
	if email == "" || password == "" {
		c.String(http.StatusBadRequest, "")
		return
	}

	user, err := db.GetUserByEmail(email)
	if err == mongo.ErrNoDocuments {
		c.String(http.StatusBadRequest, "")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.String(http.StatusUnauthorized, "")
		return
	}
	// need to perform password validation

	if err != nil {
		c.String(http.StatusInternalServerError, "")
		return
	}

	apiSecret := []byte("createapp")
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(apiSecret)
	if err != nil {
		fmt.Println("error generating token: ", err)

	}
	fmt.Println("token: ", tokenStr)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenStr,
	})

}
