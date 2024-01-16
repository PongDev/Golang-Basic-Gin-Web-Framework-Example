package service

import (
	"github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/dto"
	"github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/repository"
)

type BlogService interface {
	Create(title string, content string) error
	Get(id uint) (dto.GetBlogResponse, error)
	GetAll() (dto.GetAllBlogResponse, error)
	Update(id uint, title string, content string) error
	Delete(id uint) error
}

type blogService struct {
	r *repository.Repository
}

func newBlogService(r *repository.Repository) BlogService {
	return &blogService{r: r}
}

func (b *blogService) Create(title string, content string) error {
	return b.r.Blog.Create(title, content)
}

func (b *blogService) Get(id uint) (dto.GetBlogResponse, error) {
	blog, err := b.r.Blog.Get(id)
	if err != nil {
		return dto.GetBlogResponse{}, err
	}
	return dto.GetBlogResponse{
		ID:      blog.ID,
		Title:   blog.Title,
		Content: blog.Content,
	}, nil
}

func (b *blogService) GetAll() (dto.GetAllBlogResponse, error) {
	blogs, err := b.r.Blog.GetAll()
	if err != nil {
		return nil, err
	}
	res := make(dto.GetAllBlogResponse, len(blogs))
	for idx, blog := range blogs {
		res[idx] = dto.Blog{
			ID:      blog.ID,
			Title:   blog.Title,
			Content: blog.Content,
		}
	}
	return res, nil
}

func (b *blogService) Update(id uint, title string, content string) error {
	return b.r.Blog.Update(id, title, content)
}

func (b *blogService) Delete(id uint) error {
	return b.r.Blog.Delete(id)
}
