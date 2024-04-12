package domain

type Seller struct {
	Name string `json:"name" validate:"required,min=5"`
}