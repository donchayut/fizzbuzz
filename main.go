package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pallat/hello/fizzbuzz"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	r := gin.Default()
	r.GET("/token", tokenHandler)
	r.GET("/fizzbuzz/:number", middleware, fizzbuzzHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

var mySigningKey = []byte("AllYourBase")

// https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac
func tokenHandler(c *gin.Context) {

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		Issuer:    "pallat",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}

func wrapper(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return mySigningKey, nil
		})

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			next(c)
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
}

func middleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")[7:]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Next()
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}
}

func fizzbuzzHandler(c *gin.Context) {
	number := c.Param("number")
	n, err := strconv.Atoi(number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	type response struct {
		Number  string `json:"number" xml:"number"`
		Message string `json:"result" xml:"result"`
	}

	c.JSON(http.StatusOK, response{
		Number:  number,
		Message: fizzbuzz.Say(n),
	})
}
