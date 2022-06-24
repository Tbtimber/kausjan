package main

import (
	"net/http"

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
		re, err := ahp.ParseFrom(testInputA)
		c.JSON(http.StatusOK, gin.H{
			"response": re,
			"err":      err,
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
		}
		
	],
	"comparisons": [
		{
			"parent_id": "testA",
			"comparisons_matrix": 
			[[1,2],
			[0.5,1]]
		}
	]
}`
