package router

import (
	"github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/handler"
	"github.com/gin-gonic/gin"
)

func setupBlogRouter(r *gin.RouterGroup, h handler.BlogHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
