package models

// Book represents a book model with author details
type Book struct {
	ID       uint    `gorm:"primary_key" json:"id"`
	Title    string  `json:"title"`
	AuthorID uint    `json:"author_id"`
	Author   Author  `gorm:"foreignkey:AuthorID" json:"author"`
	Ranking  float64 `json:"ranking"`
	Price    float64 `json:"price"`
}
