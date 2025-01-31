// https://go.dev/doc/tutorial/web-service-gin

// functions need breakling up to reusable functions!!
package main

import (
	"context"
	p "entenchef/web-service-gin/pages"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mediocregopher/radix/v4"
)

func redisSearchTitle(c *gin.Context) {

	ctx := context.Background()
	client, err := (radix.PoolConfig{}).New(ctx, "tcp", "10.27.27.164:6379")
	if err != nil {
		// handle error
		println(err.Error())
	}

	var input string = c.Param("input")
	query := strings.Replace(input, "%20", " ", -1)
	query = strings.ReplaceAll(query, "\"", "")

	var res string

	errOut := client.Do(ctx, radix.Cmd(&res, "JSON.GET", query))
	println(errOut)
	println(res)
	//echo json
	c.JSON(http.StatusOK, res)

}

func redisSearchWord(c *gin.Context) {
	ctx := context.Background()
	client, err := (radix.PoolConfig{}).New(ctx, "tcp", "10.27.27.164:6379")
	if err != nil {
		// handle error
		println(err.Error())
	}

	var word string = c.Param("word")
	word = "*" + word + "*"

	var res []string
	errOut := client.Do(ctx, radix.Cmd(&res, "KEYS", word))
	println(errOut)

	c.IndentedJSON(http.StatusOK, res)

}

func main() {

	router := gin.Default()
	//gets JSON by title
	router.GET("/searched/:input", redisSearchTitle)
	//searches for word
	router.GET("/search/:word", redisSearchWord)

	router.GET("/catolog/pizza", p.RedisPizzas)
	router.Run("localhost:8080")
}
