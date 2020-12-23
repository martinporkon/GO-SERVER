package database

import (
	"database/sql"	"log"
)


var DbConn *sql.DB

func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:password123@tcp(127.0.0.1:3306)/invetorydb")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(4)
	DbConn.SetMaxIdleConns(4)
	DbConn.SetConnMazLifetime(60 * time.Second)
}

// we need the Go database driver, as this is not included in the main package.


// query the database

func getProductList() ([]Product, error) {// DATABASES MUST BE UPDATED TO USE CONTEXTS WITH TIMEOUTS.
	results, err := database.DbConn.QueryContext(`SELECT productId,
	manufacturer,
	sku,
	upc,
	pricePerUnit,
	quantityOnHand,
	productName
	FROM products`)
	if err != nil {
		return nil, err
	}
	defer results.close()
	products := make([]Product, 0)
	for results.Next() {
		var product Product
		// all must be in the same order
		results.Scan(&product.ProductID,
			 &product.Manufacturer,
			  &product.Sku,
			   &product.Upc,
			&product.PricePerUnit,
		&product.QuantityOnHand,
	&product.ProductName)
	products = append(products, product)
	}
	return products, nil
}

func getProduct(productID int) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)// if a query takes longer than 15 seconds. It will cancel and will return.
	defer cancel()
	row := database.DbConn.QueryRowContext(`SELECT productId,
	manufacturer,
	sku,
	upc,
	pricePerUnit,
	quantityOnHand,
	productName
	FROM products
	WHERE productId = ?`, productID)// those are query parameters
	product := &Product{}
	err := row.Scan(&product.ProductID,
		&product.Manufacturer,
		 &product.Sku,
		  &product.Upc,
	   		&product.PricePerUnit,
   		&product.QuantityOnHand,
		&product.ProductName)
	if err == sql.ErrNoRows {}
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return product, nil
	// w.WriteHeader(InternalServeError)
	//Works with Postman.
}

func updateProduct(product Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)// if a query takes longer than 15 seconds. It will cancel and will return.
	defer cancel()
	_, err := database.DbConn.ExecContext(`UPDATE products SET ...`)
// if err nil, return err
return nil
}

func insertProduct(product Product) (int,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)// if a query takes longer than 15 seconds. It will cancel and will return.
	defer cancel()
	result, err := database.DbConn.ExecContext(`INSERT INTO products...`)
// if err nil, return err
insertID, err := result.LastInserId()
return int(insertID), nil
}

// delete as well