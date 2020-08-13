package main

import (
	"OnlineShop/products"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
var (
	PORT string = "8000"
)

func main() {
	postgreConf := products.Config{
		Host:             "localhost",
		User:             "postgres",
		Password:         "passanya",
		Port:             5432,
		Database:         "griffon",
		ConnectionString: "",
		Params:           "sslmode=disable",
	}
	productStore, err := products.NewPostgreStore(postgreConf)
	if err != nil {
		log.Fatal(err)
	}
	productService := products.NewProductService(productStore)
	productHttpEndpoints := products.NewHttpEndpoints(productService)
	router:=mux.NewRouter()
	router.Methods("POST").Path("/products").HandlerFunc(productHttpEndpoints.CreateProductEndpoint())
	log.Println("Server is running on port "+PORT)
	err = http.ListenAndServe(":"+PORT,router)
	if err!=nil{
		log.Fatal(err)
	}
}
