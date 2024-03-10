package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]any{
			"messages": map[string]string{
				"t" : "u",
				"v" : "w",
			},
		})
	})

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("could not start the server %v", err))
	}
}
