package models

type Banner struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImgURL      string `json:"img_url"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type BannerRepo interface {
	Create(banner *Banner) error
	GetByID(id int64) (*Banner, error)
	Update(banner *Banner) error
	Delete(id int64) (*Banner, error)
	ListActive() ([]*Banner, error)
	List(page, limit int) ([]*Banner, error)
}

type BannerUsecase interface {
	Create(banner *Banner) error
	GetByID(id uint) (*Banner, error)
	Update(banner *Banner) error
	Delete(id uint) (*Banner, error)
	ListActive() ([]*Banner, error)
	List(page, limit int) ([]*Banner, error)
}
