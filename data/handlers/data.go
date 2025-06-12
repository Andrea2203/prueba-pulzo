package handlers

import (
	"net/http"
	"io"
	"fmt"
	"os"
	"log"
)

//Obtener data de la api externa
func GetData(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Printf("Método no permitido: %s\n", r.Method)
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	api := r.URL.Query().Get("api")
	apiURL := os.Getenv("RICK_MORTY_API_URL")

	if apiURL == "" || api== ""{
		apiURL = "https://rickandmortyapi.com/api"
	}
	
	if api == "" {
		api = "character"
	}

	fullURL := apiURL + "/" + api
	log.Printf("Url consultada: %s\n", fullURL)
	resp, err := http.Get(fullURL)
	
	if err != nil {
		log.Printf("Error al consultar API Externa: %s\n", err)
		http.Error(w, "Error al consultar API Externa", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		log.Printf("Respuesta de la api externa: %d\n", resp.StatusCode)
		http.Error(w, fmt.Sprintf("Respuesta de la api externa: %d\n", resp.StatusCode), http.StatusServiceUnavailable)
		return
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error al procesar respuesta de la API: %s\n", err)
		http.Error(w, "Error al procesar respuesta de la API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}