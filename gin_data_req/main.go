package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", Gethome)
	r.POST("/", Posthome)
	r.GET("/query", Querystring)         // /query?name=lol&age=21
	r.GET("/path/:name/:age", Pathparam) // /query/name/21
	r.POST("/body", Posthomebody)
	r.Run()
}

func Gethome(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Hello World",
	})
}

func Posthome(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Hello World Post",
	})
}

func Posthomebody(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"msg": string(value),
	})
}

func Querystring(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func Pathparam(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
