package handlers

import (
	"example/web-service-gin/domain"
	"example/web-service-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service services.AlbumService
}

func NewHandler(service services.AlbumService) *Handler {
	return &Handler{service}
}

func (h *Handler) GetAlbums(c *gin.Context) {
	albums, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func (h *Handler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

func (h *Handler) PostAlbums(c *gin.Context) {
	var newAlbum domain.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(newAlbum); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newAlbum)
}

func (h *Handler) DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "album deleted"})
}
