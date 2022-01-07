package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")

		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    http.StatusNotFound,
				"error_message": "路由未定义,请确认 url 和请求方法是否正确.",
			})
		}
	})

	r.Run(":8000")
}
