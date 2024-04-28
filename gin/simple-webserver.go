package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	r := gin.Default()
	r.GET("/", homePageHandler)
	return r
}
func homePageHandler(c *gin.Context) {
	c.String(200, "hello, welcome to my homepage")
}
func main() {
	err := router().Run()
	if err != nil {
		fmt.Printf("error running server: %s", err)
		return
	}
}
