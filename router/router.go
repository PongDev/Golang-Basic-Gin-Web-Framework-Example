package router

import (
	"github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, h *handler.Handler) {
	blogRouter := r.Group("/blog")
	{
		setupBlogRouter(blogRouter, h.Blog)
	}
}
