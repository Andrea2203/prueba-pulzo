package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"pulzo/auth/handlers"
)

func main() {
	http.HandleFunc("/create-token", handlers.GenerateToken)
	http.HandleFunc("/use-api", handlers.UseToken)

	fmt.Println("Servidor en http://localhost:8080")
	port := os.Getenv("AUTH_PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}