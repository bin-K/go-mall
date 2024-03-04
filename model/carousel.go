package model

import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	ImagePath string
	ProductId int `gorm:"not null"`
}
