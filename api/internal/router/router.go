package router

import (
	"gin-api/internal/handler"
	"gin-api/internal/pkg/ginautowrap"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouter(g *gin.Engine) {
	g.GET("/api/status/ready", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "ready"})
	})
	g.GET("/api/addone", ginautowrap.AutoBindWrap(handler.AddOne))
	g.GET("/api/map", ginautowrap.AutoBindWrap(handler.GetMap))
	g.GET("/api/slice", ginautowrap.AutoBindWrap(handler.GetSlice))
	g.GET("/api/float", ginautowrap.AutoBindWrap(handler.GetFloat))
	g.GET("/api/interface", ginautowrap.AutoBindWrap(handler.GetInterface))

	demoUser := g.Group("/api/demo")
	{
		demoUser.POST("/user", ginautowrap.AutoBindWrap(handler.AddUser))
		demoUser.DELETE("/user", ginautowrap.AutoBindWrap(handler.DelUser))
	}
}
