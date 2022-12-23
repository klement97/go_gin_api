package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"italgold/controller"
	"italgold/database"
	"italgold/middleware"
	"italgold/model"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{}, &model.Entry{})
	serveApp()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env.local file")
	}
}

func serveApp() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRouter := router.Group("/api")
	protectedRouter.Use(middleware.JWTAuthMiddleware())
	protectedRouter.POST("/entry", controller.AddEntry)
	protectedRouter.GET("/entry", controller.GetAllEntries)

	router.Run()
	fmt.Println("Server started on port 8000")
}
