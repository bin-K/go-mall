package dao

import (
	"fmt"
	"go-mall/model"
)

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
		&model.Admin{},
		&model.Order{},
		&model.Category{},
		&model.Address{},
		&model.Carousel{},
		&model.Cart{},
		&model.Favorite{},
		&model.Notice{},
		&model.Product{},
		&model.ProductImg{},
	)
	if err != nil {
		fmt.Println(err, "Migration err")
	}
	return
}
