package models

import "time"

type Tour struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Duration    int       `json:"duration"`
	HotelID     *Hotel    `json:"hotelID"`
	Images      []string  `json:"images"`
	StartDate   []string  `json:"startDate"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TourRepo interface {
	Create(tour *Tour) error
	Update(tour *Tour) error
	Delete(id uint) error
	GetById(id uint) (*Tour, error)
	List(page, limit int) ([]*Tour, error)
}

type TourUsecase interface {
	Create(tour *Tour) error
	Update(tour *Tour) error
	Delete(id uint) error
	GetById(id uint) (*Tour, error)
	List(page, limit int) ([]*Tour, error)
}
