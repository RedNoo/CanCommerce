package services

import (
	models "CanCommerce/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

func productList(db *gorm.DB) ([]*models.Product, error) {
	fmt.Println("productList services")

	pd := &models.ProductDB{}
	pd.Db = db

	ll, err := pd.List()

	if err != nil {
		panic(err.Error())
	}

	return ll, err
}
