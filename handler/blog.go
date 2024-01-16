package handler

import (
	"net/http"
	"strconv"

	"github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/dto"
	"github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/service"
	"github.com/gin-gonic/gin"
)

type BlogHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type blogHandler struct {
	blogService service.BlogService
}

func newBlogHandler(s service.BlogService) BlogHandler {
	return &blogHandler{blogService: s}
}

func (b *blogHandler) Create(c *gin.Context) {
	req := dto.CreateBlogRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := b.blogService.Create(req.Title, req.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Create a Blog Successfully."})
}

func (b *blogHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blog, err := b.blogService.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blog)

}

func (b *blogHandler) GetAll(c *gin.Context) {
	blogs, err := b.blogService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

func (b *blogHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req := dto.UpdateBlogRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = b.blogService.Update(uint(id), req.Title, req.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update a Blog Successfully."})
}

func (b *blogHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = b.blogService.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete a Blog Successfully."})
}
