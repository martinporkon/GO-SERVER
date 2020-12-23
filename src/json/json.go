package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	ProductID      int    `json:"productId`
	Manufacturer   string `json:"manufacturer`
	Sku            string `json:"sku`
	Upc            string `json:"upc`
	PricePerUnit   string `json:"pricePerUnit`
	QuantityOnHand int    `json:"quantityOnHand`
	ProductName    string `json:"productName`
}

var productList []Product

func init() {

}

func main() {

	product := &Product{
		ProductID:      123,
		Manufacturer:   "Big Box Company",
		PricePerUnit:   "12.99",
		Sku:            "4561qHJK",
		Upc:            "414654444566",
		QuantityOnHand: 28,
		ProductName:    "Gizmo",
	}
	productJSON, err := json.Marshal(product)
	// this will return a byte slice or a possible error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(productJSON)) // prints out the JSON
	fmt.Println("Hello, playground")
	product2 := Product{}
	err2 := json.Unmarshal([]byte(productJSON), &product2)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("json was unmarshalled product: %s", product.ProductName)
}
