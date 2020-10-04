package shorten

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
	ID       string `json:"id"`
	Location string `json:"location"`
}
