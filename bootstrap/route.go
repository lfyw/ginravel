package bootstrap

import (
	"ginravel/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {

	registerGlobalMiddleware(router)

	routes.RegisterApiRoutes(router)

	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面未找到")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    http.StatusNotFound,
				"error_message": "路由未定义,请确认 url 和请求方法是否正确.",
			})
		}
	})
}
