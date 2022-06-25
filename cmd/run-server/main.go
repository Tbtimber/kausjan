package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Tbtimber/kausjan/ahp"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ahp/tree", func(c *gin.Context) {
		start := time.Now()
		re, err := ahp.ParseFrom(testInputA)
		re2, err2 := ahp.ConvertInputToAhpTree(re)
		ahp.ComputeTree(&re2)
		fmt.Println(re2)
		c.IndentedJSON(http.StatusOK, gin.H{
			"response":    re,
			"err":         err,
			"response2":   re2,
			"err2":        err2,
			"timeElapsed": time.Since(start),
		})
	})
	r.Run()

}

var testInputA = `
{
	"leaf_array" : [
		{
			"id": "testA",
			"parent_id": ""
		},
		{
			"id": "testB",
			"parent_id": "testA"
		},
		{
			"id": "testC",
			"parent_id": "testA"
		},
		{
			"id": "testD",
			"parent_id": "testB"
		},
		{
			"id": "testE",
			"parent_id": "testB"
		}
		
	],
	"comparisons": {
		"testA": {
			"order" : ["testB", "testC"],
			"comparison_matrix": 
			[[1,2],
			[0.5,1]]
		},
		"testB": {
			"order" : ["testD", "testE"],
			"comparison_matrix": 
			[[1,2],
			[0.5,1]]
		}
	}
}`
