package handler

import "github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/service"

type Handler struct {
	Blog BlogHandler
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		Blog: newBlogHandler(s.Blog),
	}
}
