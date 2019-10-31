package services

import (
	models "CanCommerce/models"
	services "CanCommerce/services"
	"gologi/services/httpUtils"
	"net/http"
)

type ReturnAddProduct struct {
	ID string `json:"id"`
}

func AddProduct(w http.ResponseWriter, p models.Product) {
	product_id, err := services.AddProduct(p.SKU, p.Title, p.OldPrice, p.Price, p.ShortDescription, p.LongDescription, p.Tag, p.Quantity, true, p.IsSale, p.IsActive)
	if err != nil {
		httpUtils.HandleError(&w, 200, err.Error(), "", nil)
	}

	httpUtils.HandleSuccess(&w, ReturnAddProduct{ID: string(product_id)})
}

func ListProducts(w http.ResponseWriter, page int) {
	list := services.GetProducts(page)

	httpUtils.HandleSuccess(&w, list)
}
