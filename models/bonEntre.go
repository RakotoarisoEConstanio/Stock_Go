package models

import (
	"time"
)

// BonEntree représente une livraison ou une entrée de matériaux
type BonEntree struct {
	ID            uint        `gorm:"primaryKey" json:"id"`
	Date          time.Time   `gorm:"not null" json:"date"`
	MateriauID    uint        `gorm:"not null" json:"materiau_id"`
	Materiaux     Materiaux   `gorm:"foreignKey:MateriauID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"materiaux"`
	Quantite      float64     `gorm:"not null" json:"quantite"`
	FournisseurID uint        `gorm:"not null" json:"fournisseur_id"`
	Fournisseur   Fournisseur `gorm:"foreignKey:FournisseurID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"fournisseur"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
