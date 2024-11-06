package main

import (
	"example/web-service-gin/infra"
	"example/web-service-gin/internal/albums/handlers"
	"example/web-service-gin/internal/albums/repositories"
	"example/web-service-gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.InitDB()
	repo := repositories.NewRepository(infra.DB)
	service := services.NewAlbumService(repo)
	handler := handlers.NewHandler(service)

	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumByID)
	router.POST("/albums", handler.PostAlbums)
	router.DELETE("/albums/:id", handler.DeleteAlbumByID)

	router.Run(":8080")
}
