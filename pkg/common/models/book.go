package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model  // gorm.Model will add properties such as ID, CreatedAt, UpdatedAt and DeletedAt for us.
	Title       string
	AuthorId    int
	Description string
}
