package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type fooHandler struct {
	Message string
}

type Product struct {
	ProductID      int    `json:"productId`
	Manufacturer   string `json:"manufacturer`
	Sku            string `json:"sku`
	Upc            string `json:"upc`
	PricePerUnit   string `json:"pricePerUnit`
	QuantityOnHand int    `json:"quantityOnHand`
	ProductName    string `json:"productName`
}

// in order to have at this interface we need a ServeHttp func handler

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message)) //<< wrties out the message field to the HTTP response using the message writer
}

func main() { // see fail on main fail serveri alustamiseks ning tööle panemiseks.
	http.Handle("/foo", &fooHandler{Message: "foo called"}) // set the message
	// set the serer to listen and serve
	http.HandleFunc("/bar", barHandler) // bar patterna and the HTTP handleFunc funciton
	http.ListenAndServe(":5000", nil)   // nil for the handler and ServeMux. This will tell it to use the default ServeMux

	http.HandleFunc("/products", productHandler)

}

// here is the simpler Http call function

func barHandler(w http.ResponseWriter, r *http.Request) {
	// accepts a response writer and a pointer to the request
	w.Write([]byte("bar called"))
}

func getNextID() int {
	highestID := -1
	for _, product := range productList {
		if highestID < product.ProductID {
			highestID = product.ProductID
		}
	}
	return highestID + 1
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	// handlers are capable of handling request messages with different request methods.
	switch r.Method {
	case http.MethodGet:
		productJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJson)
	case http.MethodPost:
		// add a new product to the list
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body) // to read out the bytes to the memory
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return // HTTP 400 status.
		}
		newProduct.ProductID = getNextID()
		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

// ServeMUX will find the correct variable tpe patameters

// Dybamic or Parametric Routes
// /products/123
