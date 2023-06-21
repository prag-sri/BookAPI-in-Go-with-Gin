package models

type Author struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	EmailID string
	Age     uint
	Country string
	Books   []Book // Association with Book model
}
