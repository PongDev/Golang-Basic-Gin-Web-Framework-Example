package service

import "github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/repository"

type Service struct {
	Blog BlogService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Blog: newBlogService(r),
	}
}
