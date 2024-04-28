package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	r := gin.Default()
	userRoute := r.Group("/user")
	{
		userRoute.GET("/hello/:name", Handler)
	}
	return r
}
func Handler(c *gin.Context) {
	user := c.Param("name")
	c.String(200, fmt.Sprintf("hello %s, welcome to my homepage", user))
}
func main() {
	err := router().Run()
	if err != nil {
		fmt.Printf("error running server: %s", err)
		return
	}
}
