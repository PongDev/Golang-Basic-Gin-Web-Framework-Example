package repository

import (
	"fmt"

	"github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	Blog BlogRepository
}

func NewRepository(host, user, password, dbname string, port int, timezone string) *Repository {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		host, user, password, dbname, port, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Blog{})
	return &Repository{
		Blog: newBlogRepository(db),
	}
}
