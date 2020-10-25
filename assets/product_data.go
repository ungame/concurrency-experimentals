package assets

import (
	"concurrency-experimentals/models"
	"concurrency-experimentals/utils"
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

type ProductData struct {
	Name         string     `json:"name"`
	Price        float64    `json:"price"`
	Category     []Category `json:"category"`
	Description  string     `json:"description"`
	ImageURL     string     `json:"image"`
	Manufacturer string     `json:"manufacturer"`
}

type Category struct {
	Name string `json:"name"`
}

func (p *ProductData) GenerateProduct(r *rand.Rand) *models.Product {
	product := new(models.Product)
	product.Name = p.Name
	product.Price = p.Price
	product.Description = p.Description
	product.PhotoURL = p.ImageURL
	product.Manufacturer = p.Manufacturer
	if len(p.Category) > 0 {
		product.Category = p.Category[0].Name
	}
	product.GenerateID()
	product.GenerateQuantity(r)
	product.Ratings = r.Intn(1000)
	product.GenerateCreatedAt(r)
	product.GenerateUpdatedAt(r)
	product.GenerateDeletedAt(r)
	return product
}

func LoadProductData() []ProductData {
	dirName, _ := os.Getwd()
	b := utils.LoadFromFile(dirName + "/assets/products.json")
	data := []ProductData{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		log.Println("Unmarshal Error: ", err.Error())
		return nil
	}
	return data
}
