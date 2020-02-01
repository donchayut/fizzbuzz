package main

import (
	"encoding/xml"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pallat/hello/fizzbuzz"
)

func main() {
	r := gin.Default()
	r.GET("/fizzbuzz/:number", fizzbuzzHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
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
		XMLName xml.Name `xml:"Envelop"`
		Number  string   `json:"number" xml:"number"`
		Message string   `json:"result" xml:"result"`
	}

	c.XML(http.StatusOK, response{
		Number:  number,
		Message: fizzbuzz.Say(n),
	})
}
