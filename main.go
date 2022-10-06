package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/:name", greet)
	r.Run()
}

func greet(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}
