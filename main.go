package main

import (
	"database/sql"
	"demo/database"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint 1\n")
	fmt.Fprint(w, "Next endpoint: /page2")
}

func page2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Homepage Endpoint 2")

	query := "SELECT * FROM product WHERE productId = 001"
	result, err := Globaldb.Query(query)
	if err != nil {
		log.Fatal("Error Query!")
	}
	defer result.Close()

	if !result.Next() {
		fmt.Fprint(w, "No data with productId 001")
		return
	}
	product := database.Product{}
	err = result.Scan(&product.ProductId, &product.ProductName, &product.ProductPrice, &product.ProductDescription)

	if err != nil {
		log.Fatal("Failed!")
	}
	fmt.Fprintf(w, "ProductId: %d | ProductName: %s | ProductPrice: %d | ProductDescription: %s", product.ProductId, product.ProductName, product.ProductPrice, product.ProductDescription)
}

func handleRequests() {
	log.Print("Server Running at: http://localhost:8080")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/page2", page2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Fail to run server")
	}
}

var Globaldb *sql.DB

func main() {
	var err error
	// Globaldb, err = sql.Open("mysql", "root:purnama9@tcp(localhost:3306)/piplupShop")
	Globaldb, err = sql.Open("mysql", "root:purnama9@tcp(database-1.cnq6mumoop1p.us-east-1.rds.amazonaws.com:3306)/piplupShop")
	if err != nil {
		log.Fatal("Fail to connect to database 1")
	}
	if Globaldb.Ping() != nil {
		log.Fatal("Fail to connnect to database 2")
	}
	log.Println("Connected to database")

	handleRequests()
}
