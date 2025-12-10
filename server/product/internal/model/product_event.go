package model

type ProductEvent struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
	Slug string `json:"slug"`
}
