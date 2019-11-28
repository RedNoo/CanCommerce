package services

import (
	dbUtil "CanCommerce/db"
	"fmt"
	"log"
	"time"

	models "CanCommerce/models"
)

//AddProduct : Adding product
func AddProduct(sku string, title string, old_price float64, price float64, short_description string, long_description string, tag string, quantity int, is_new bool, is_sale bool, is_active bool) (int64, error) {

	var lastInsertId int64

	currentTime := time.Now()
	currentTime.Format("2006-01-02 15:04:05.000000")
	err := dbUtil.Db_.QueryRow("INSERT INTO products( sku, title, old_price, price, short_description, long_description, tag, quantity, is_new, is_sale, is_active, created_at, deleted_at, updated_at) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) returning id;", sku, title, old_price, price, short_description, long_description, tag, quantity, is_new, is_sale, is_active, currentTime, nil, nil).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	fmt.Println(lastInsertId)
	return lastInsertId, err

}

func GetProductCount() int {
	statement := `SELECT  count(*)	FROM products;`
	rows, err := dbUtil.Db_.Query(statement)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for rows.Next() {

		rows.Scan(&count)
	}

	return count
}

// GetProducts : Listing all products
func GetProducts(page int) []models.Product {

	statement := `SELECT  id, sku, title, old_price, price, short_description, long_description, tag, quantity, is_new, is_sale, is_active, created_at, deleted_at, updated_at 	FROM products  limit 20 offset $1;`
	rows, err := dbUtil.Db_.Query(statement, ((page - 1) * 20))
	if err != nil {
		log.Fatal(err)
	}

	productModelList := make([]models.Product, 0)

	for rows.Next() {
		p := new(models.Product)

		err := rows.Scan(&p.ID, &p.SKU, &p.Title, &p.OldPrice, &p.Price, &p.ShortDescription, &p.LongDescription, &p.Tag, &p.Quantity, &p.IsNew, &p.IsSale, &p.IsActive, &p.CreatedAt, &p.DeletedAt, &p.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}

		productModelList = append(productModelList, *p)
	}

	defer rows.Close()

	return productModelList
}
