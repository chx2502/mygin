package main

import (
	//"fmt"
	"net/http"
	"mygin"
)

func main() {
	r := mygin.New()
	r.GET("/", func(c *mygin.Context) {
		c.HTML(http.StatusOK, "<h1>Home Page</h1>")
	})

	r.GET("/hello", func(c *mygin.Context) {
		c.String(http.StatusOK, "hello %s, you are at: %s\n", c.Query("name"), c.Path)
	})

	r.GET("hello/:name", func(c *mygin.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *mygin.Context) {
		c.JSON(http.StatusOK, mygin.H{"filepath": c.Param("filepath")})
	})

	r.POST("/login", func(c *mygin.Context) {
		c.JSON(http.StatusOK, mygin.H {
			"username": c.PostFrom("username"),
			"password": c.PostFrom("password"),
		})
	})

	r.Run(":9999")
}
