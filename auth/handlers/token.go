package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"errors"
	"fmt"
    "io"
    "net/url"
    "os"
	"log"

    "pulzo/auth/models"
	"github.com/google/uuid"
)
//Almacenamiento de tokens
var (
	tokenStore = make(map[uuid.UUID]*models.Token)
	mutex      = sync.Mutex{}
)

//Generar Token
func GenerateToken(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		log.Printf("Método no permitido: %s\n", r.Method)
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	
	token := &models.Token{
		Value:   uuid.New(),
		Uses:    5,
		IsValid: true,
	}

	mutex.Lock()
	tokenStore[token.Value] = token
	mutex.Unlock()
	log.Printf("Token generado: %s\n", token.Value.String())
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

//UseToken verifica el token obtenido y realiza la solicitud
func UseToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("Método no permitido: %s\n", r.Method)
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	tokenID := r.URL.Query().Get("token")
	if tokenID == "" {
		log.Println("El Token no ha sido enviado")
		http.Error(w, "El Token no ha sido enviado", http.StatusBadRequest)
		return
	}
	token, err := ValidToken(tokenID)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusGone)
		return
	}
	api := r.URL.Query().Get("api")
	if api == "" {
		api = "character"
	}
	data, err := getData(api)
	if err != nil {
		log.Printf(err.Error())
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	response := map[string]interface{}{
		"token": token,
		"data":  json.RawMessage(data),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

//ValidToken valida si el token es valido
func ValidToken(tokenID string) (*models.Token, error){
	id, err := uuid.Parse(tokenID)
	if err != nil {
		return nil, errors.New("El token no es válido")
	}

	mutex.Lock()
	defer mutex.Unlock()

	token, exists := tokenStore[id]
	if !exists {
		return nil, errors.New("El token no existe")
	}

	if !token.IsValid || token.Uses <= 0 {
		token.IsValid = false
		return nil, errors.New("El token no es valido")
	}

	token.Uses--
	if token.Uses == 0 {
		token.IsValid = false
	}
	
	log.Printf("usos restantes del token: %d\n", token.Uses)
	return token, nil

}

//getData realiza la peticion y obtiene la información en formato JSON
func getData(api string) ([]byte, error) {
	
	dataURL := os.Getenv("DATA_URL")
	if dataURL == "" {
		dataURL = "http://localhost:8081/get-data"
	}
	
	fullURL := fmt.Sprintf("%s?api=%s", dataURL, url.QueryEscape(api))
	log.Printf("Url consultada: %s\n", fullURL)
	
	resp, err := http.Get(fullURL)
	
	if err != nil {
		log.Printf("Error al obtener datos: %s\n", err)
		return nil, fmt.Errorf("Error al obtener datos: %s\n", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		log.Printf("Respuesta de la api externa: %d\n", resp.StatusCode)
		return nil, fmt.Errorf("Respuesta de la api externa: %d\n", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error al leer la informacíon: %s\n", err)
		return nil, fmt.Errorf("Error al leer la informacíon: %s\n", err)
	}

	return body, nil

}