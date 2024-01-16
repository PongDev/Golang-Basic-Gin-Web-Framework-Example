package repository

import (
	"github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/model"
	"gorm.io/gorm"
)

type BlogRepository interface {
	Create(title string, content string) error
	Get(id uint) (*model.Blog, error)
	GetAll() ([]model.Blog, error)
	Update(id uint, title string, content string) error
	Delete(id uint) error
}

type blogRepository struct {
	db *gorm.DB
}

func newBlogRepository(db *gorm.DB) BlogRepository {
	return &blogRepository{db: db}
}

func (b *blogRepository) Create(title string, content string) error {
	blog := model.Blog{
		Title:   title,
		Content: content,
	}
	result := b.db.Create(&blog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *blogRepository) Get(id uint) (*model.Blog, error) {
	blog := model.Blog{}
	result := b.db.First(&blog, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &blog, nil
}

func (b *blogRepository) GetAll() ([]model.Blog, error) {
	blogs := []model.Blog{}
	result := b.db.Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogs, nil
}

func (b *blogRepository) Update(id uint, title string, content string) error {
	blog := model.Blog{}
	result := b.db.First(&blog, id)
	if result.Error != nil {
		return result.Error
	}

	blog.Title = title
	blog.Content = content
	result = b.db.Save(&blog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *blogRepository) Delete(id uint) error {
	result := b.db.Delete(&model.Blog{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
