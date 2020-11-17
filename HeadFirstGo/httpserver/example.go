package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	r.POST("/post", postBodyProcess)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func postBodyProcess(c *gin.Context) {

	type member struct {
		Name string `json:"name"`
		Job  string `json:"job"`
		Age  string `json:"age"`
	}

	var person member

	c.ShouldBind(&person)

	person.Age = "30"

	// fmt.Printf("obj %s:", obj)

	c.JSON(200, person)
}
