package repository

// LAYER 1: REPOSITORY (Akses Database via GORM)

type HelloRepository interface {
	GetMessage() string
}

type helloRepository struct{}

func NewHelloRepository() HelloRepository {
	return &helloRepository{}
}

func (r *helloRepository) GetMessage() string {
	return "Hello World!"
}
