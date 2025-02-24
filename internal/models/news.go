package models

import "time"

type New struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"imageURL"`
	CreatedAt time.Time `json:"createdAt"`
	Updated   time.Time `json:"updated"`
}

type NewRepo interface {
	Create(new *New) error
	Update(new *New) error
	Delete(id uint) error
	GetById(id uint) (*New, error)
	List(page, limit int) ([]*New, error)
}

type NewUsecase interface {
	Create(new *New) error
	Update(new *New) error
	Delete(id uint) error
	GetById(id uint) (*New, error)
	List(page, limit int) ([]*New, error)
}
