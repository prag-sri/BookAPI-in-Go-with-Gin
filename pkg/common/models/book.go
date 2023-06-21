package models

type Book struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	AuthorID    uint   // Foreign key for the associated author
	Author      Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Association with Author model
}
