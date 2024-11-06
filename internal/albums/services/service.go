package services

import (
	"example/web-service-gin/internal/albums/domain"
	"example/web-service-gin/internal/albums/repositories"
)

type AlbumService interface {
	GetAll() ([]domain.Album, error)
	GetByID(id string) (domain.Album, error)
	Create(album domain.Album) error
	DeleteByID(id string) error
}

type albumService struct {
	repo repositories.Repository
}

func NewAlbumService(repo repositories.Repository) AlbumService {
	return &albumService{repo}
}

func (s *albumService) GetAll() ([]domain.Album, error) {
	return s.repo.GetAll()
}

func (s *albumService) GetByID(id string) (domain.Album, error) {
	return s.repo.GetByID(id)
}

func (s *albumService) Create(album domain.Album) error {
	return s.repo.Create(album)
}

func (s *albumService) DeleteByID(id string) error {
	return s.repo.DeleteByID(id)
}
