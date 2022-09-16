package pizza

import (
	"context"
	//"fmt"

	//"encoding/json"
	"net/http"

	//"strings"
	"github.com/gin-gonic/gin"
	"github.com/mediocregopher/radix/v4"
)

// --------

//---------

func JSONArray(res []string) []string {
	//array to return
	out := []string{}
	ctx := context.Background()

	//redis connection
	client, err := (radix.PoolConfig{}).New(ctx, "tcp", "10.27.27.164:6379")
	if err != nil {
		// handle error
		println(err)
	}

	//append each json to array
	for i := 0; i < len(res); i++ {
		var hold string //to append
		_er := client.Do(ctx, radix.Cmd(&hold, "JSON.GET", res[i]))
		if _er != nil {
			println(_er)
		}
		/*
			convert hold to json format and out change to a form of array that holds json


		*/

		out = append(out, hold)
		println(hold, "\n")
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
