package models

import "time"

type Hotel struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Stars       int       `json:"stars"`
	Address     string    `json:"address"`
	Images      []string  `json:"images"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type HotelRepo interface {
	Create(hotel *Hotel) error
	Update(hotel *Hotel) error
	Delete(id uint) error
	GetById(id uint) (*Hotel, error)
	List(page, limit int) (*Hotel, error)
}

type HotelUsecase interface {
	Create(hotel *Hotel) error
	Update(hotel *Hotel) error
	Delete(id uint) error
	GetById(id uint) (*Hotel, error)
	List(page, limit int) ([]*Hotel, error)
}
