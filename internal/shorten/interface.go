package shorten

import "time"

// Repository ...
type Repository interface {
	Close()
	FindByID(string) (*ShortURLModel, error)
	Find() ([]*ShortURLModel, error)
	Create(*ShortURLModel) error
	Update(*ShortURLModel) error
	Delete(string) error
}

// ShortURLModel ...
type ShortURLModel struct {
	ID       string    `json:"id,omitempty"`
	Location string    `json:"location"`
	ExpireAt time.Time `json:"expire_at"`
}
