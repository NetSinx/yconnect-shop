package model

import "time"

type CategoryMirrorResponse struct {
	ID        uint      `json:"id"`
	Nama      string    `json:"nama"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
