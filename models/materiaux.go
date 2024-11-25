package models

import (
	"time"
)

// Matériaux représente un matériau dans le stock
type Materiaux struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Nom          string    `gorm:"size:100;not null" json:"nom"`
	Description  string    `gorm:"type:text" json:"description"`
	StockInitial float64   `gorm:"not null" json:"stock_initial"`
	StockActuel  float64   `gorm:"not null" json:"stock_actuel"`
	SeuilMin     float64   `gorm:"not null" json:"seuil_min"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
