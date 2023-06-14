package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model         // gorm.Model will add properties such as ID, CreatedAt, UpdatedAt and DeletedAt for us.
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`

	//The json tag at the end gives GORM the information of each column's names in our Postgres database.
}
