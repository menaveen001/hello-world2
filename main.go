package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func main() {

// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	routering()
// 	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

// }
func main() {
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("num1 ,num2")
		sum := (num1 + num2)
		c.String(http.StatusOK, "Hello %s", name, sum)
	})

	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action"
		c.String(http.StatusOK, "%t", b)

	})

	router.GET("/user/groups", func(c *gin.Context) {

		c.String(http.StatusOK, "The available groups are[..]")
	})

	router.Run(":8080")
}
