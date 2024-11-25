package controllers

import (
	"fmt"
	"symrise/database"
	"symrise/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AddEntree(c *fiber.Ctx) error {
	var data map[string]interface{}

	// Parser la requête JSON
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erreur de lecture du corps de la requête",
		})
	}

	// Vérification des champs requis
	requiredFields := []string{"date", "materiau_id", "quantite", "fournisseur_id"}
	for _, field := range requiredFields {
		if _, exists := data[field]; !exists {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("Le champ '%s' est requis", field),
			})
		}
	}

	// Conversion des champs
	materiauID, ok := data["materiau_id"].(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Le champ 'materiau_id' doit être un nombre valide",
		})
	}
	dateStr, ok := data["date"].(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Format de date invalide",
		})
	}

	dateEx, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Format de date invalide",
		})
	}

	quantite, ok := data["quantite"].(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "La quantité doit être un nombre valide",
		})
	}

	// Récupérer le matériau
	var materiau models.Materiaux
	if err := database.DB.First(&materiau, materiauID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Matériau non trouvé",
		})
	}

	// Mise à jour du stock
	materiau.StockActuel += quantite

	// Créer le bon d'entrée
	bonEntree := models.BonEntree{
		Date:          dateEx,
		MateriauID:    uint(materiauID),
		Quantite:      quantite,
		FournisseurID: uint(data["fournisseur_id"].(float64)),
	}

	// Sauvegarder le bon d'entrée et mettre à jour le stock
	if err := database.DB.Create(&bonEntree).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erreur lors de l'ajout du Bon Entree",
		})
	}

	if err := database.DB.Save(&materiau).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erreur lors de la mise à jour du stock du matériau",
		})
	}

	// Réponse de succès
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Bon Entree ajouté avec succès et stock mis à jour",
		"data":    bonEntree,
	})
}
