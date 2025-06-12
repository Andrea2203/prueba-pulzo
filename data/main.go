package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pulzo/data/handlers"
)

func main() {
	http.HandleFunc("/get-data", handlers.GetData)

	fmt.Println("Servidor en http://localhost:8081")
	port := os.Getenv("DATA_PORT")
	if port == "" {
		port = "8081"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}