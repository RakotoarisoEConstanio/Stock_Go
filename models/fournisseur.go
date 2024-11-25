package models

import (
	"time"
)

// Fournisseur représente un fournisseur de matériaux
type Fournisseur struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nom       string    `gorm:"size:100;not null" json:"nom"`
	Contact   string    `gorm:"size:100" json:"contact"`
	Email     string    `gorm:"size:100" json:"email"`
	Adresse   string    `gorm:"type:text" json:"adresse"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
