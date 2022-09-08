package pizza

import (
	"context"
	//"encoding/json"
	"net/http"

	//"strings"
	"github.com/gin-gonic/gin"
	"github.com/mediocregopher/radix/v4"
)

func Quotes(naked string) string {
	naked = "\"" + naked + "\""
	return naked
}

//--------

func JSONArray(res []string) []string {
	//array to return
	var out []string
	ctx := context.Background()

	//redis connection
	client, err := (radix.PoolConfig{}).New(ctx, "tcp", "10.27.27.164:6379")
	if err != nil {
		// handle error
		println(err)
	}

	//append each json to array
	for i := 0; i < len(res); i++ {
		var temp string = Quotes(res[i])
		var hold string //to append
		println(temp)
		println(hold)
		client.Do(ctx, radix.Cmd(&hold, "JSON.GET", temp))
	}
	return out

	//client.Do(ctx, radix.Cmd(&temp, "KEYS", get))
}

// ---------------
func RedisPizzas(c *gin.Context) {
	ctx := context.Background()
	client, err := (radix.PoolConfig{}).New(ctx, "tcp", "10.27.27.164:6379")
	if err != nil {
		// handle error
		println(err)
	}
	var res []string

	errOut := client.Do(ctx, radix.Cmd(&res, "KEYS", "*Pizza*"))
	println(errOut)
	//func for all titles grab json
	var result []string = JSONArray(res)

	c.IndentedJSON(http.StatusOK, result)
}
