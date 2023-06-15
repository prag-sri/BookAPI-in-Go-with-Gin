package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model        // gorm.Model will add properties such as ID, CreatedAt, UpdatedAt and DeletedAt for us.
	Title      string `json:"title"`
	AuthorId   int    `gorm:"column:author_id"`
	Author     Author `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	//By changing Author from Author to *Author, you're telling GORM to load the associated author as a pointer to the Author struct.
	Description string `json:"description"`

	//The json tag at the end gives GORM the information of each column's names in our Postgres database.
}
