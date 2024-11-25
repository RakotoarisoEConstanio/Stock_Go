package models

import (
	"time"
)

// BonSortie représente une sortie de matériaux
type BonSortie struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Date         time.Time `gorm:"not null" json:"date"`
	MateriauID   uint      `gorm:"not null" json:"materiau_id"`
	Materiaux    Materiaux `gorm:"foreignKey:MateriauID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"materiaux"`
	Quantite     float64   `gorm:"not null" json:"quantite"`
	Destinataire string    `gorm:"size:100;not null" json:"destinataire"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
