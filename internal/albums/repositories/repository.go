package repositories

import (
	"database/sql"
	"errors"
	"example/web-service-gin/domain"
)

type Repository interface {
	GetAll() ([]domain.Album, error)
	GetByID(id string) (domain.Album, error)
	Create(album domain.Album) error
	DeleteByID(id string) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]domain.Album, error) {
	rows, err := r.db.Query("SELECT id, title, artist, price FROM album")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []domain.Album
	for rows.Next() {
		var alb domain.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, err
		}
		albums = append(albums, alb)
	}

	return albums, nil
}

func (r *repository) GetByID(id string) (domain.Album, error) {
	var alb domain.Album
	err := r.db.QueryRow("SELECT id, title, artist, price FROM album WHERE id = ?", id).Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
	if err == sql.ErrNoRows {
		return alb, errors.New("album not found")
	}
	if err != nil {
		return alb, err
	}
	return alb, nil
}

func (r *repository) Create(album domain.Album) error {
	_, err := r.db.Exec("INSERT INTO album (id, title, artist, price) VALUES (?, ?, ?, ?)",
		album.ID, album.Title, album.Artist, album.Price)
	return err
}

func (r *repository) DeleteByID(id string) error {
	_, err := r.db.Exec("DELETE FROM album WHERE id = ?", id)
	return err
}
