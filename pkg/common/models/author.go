package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model        //gorm.model will add properties such as Id, CreatedAt, UpdatedAt and DeletedAt for us
	Name       string `json:"name"`
	EmailId    string `json:"emailId"`
	Age        int    `json:"age"`
	Country    string `json:"country"`
}
