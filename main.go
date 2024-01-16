package main

import (
	"os"
	"strconv"

	"github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/handler"
	"github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/repository"
	"github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/router"
	"github.com/PongDev/Golang-Basic-Gin-Web-Framework-Example/service"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	debugMode, err := strconv.ParseBool(os.Getenv("DEBUG_MODE"))
	if err != nil {
		panic(err)
	}
	if debugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	repo := repository.NewRepository(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), dbPort, os.Getenv("TIMEZONE"))
	s := service.NewService(repo)
	h := handler.NewHandler(s)
	router.SetupRouter(r, h)
	r.Run() // listen and serve on 0.0.0.0:8080 or 0.0.0.0:PORT if PORT env is specified.
}
