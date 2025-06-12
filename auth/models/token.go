package models

import "github.com/google/uuid"

type Token struct {
	Value   uuid.UUID `json:"value"`
	Uses  	int    `json:"uses"`
	IsValid bool   `json:"is_valid"`
}
