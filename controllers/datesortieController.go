package controllers

import (
	"fmt"
	"symrise/database"
	"symrise/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AddSortie(c *fiber.Ctx) error {
	var data map[string]interface{}

	// Parser la requête JSON
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erreur de lecture du corps de la requête",
		})
	}

	// Vérification des champs requis
	requiredFields := []string{"date", "materiau_id", "quantite", "destinataire"}
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
			"error": "Le champ 'date' doit être une chaîne valide",
		})
	}

	dateSortie, err := time.Parse("2006-01-02", dateStr)
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

	destinataire, ok := data["destinataire"].(string)
	if !ok || destinataire == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Le champ 'destinataire' est requis",
		})
	}

	// Récupérer le matériau
	var materiau models.Materiaux
	if err := database.DB.First(&materiau, uint(materiauID)).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Matériau non trouvé",
		})
	}

	// Vérification de la quantité disponible
	if materiau.StockActuel < quantite {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Stock insuffisant pour effectuer cette sortie",
		})
	}

	// Mise à jour du stock
	materiau.StockActuel -= quantite

	// Créer le bon de sortie
	bonSortie := models.BonSortie{
		Date:         dateSortie,
		MateriauID:   uint(materiauID),
		Quantite:     quantite,
		Destinataire: destinataire,
	}

	// Sauvegarder le bon de sortie et mettre à jour le stock
	if err := database.DB.Create(&bonSortie).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erreur lors de l'ajout du Bon Sortie",
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
		"message": "Bon de sortie ajouté avec succès et stock mis à jour",
		"data":    bonSortie,
	})
}

func GetBonsortie(c *fiber.Ctx) error {
	var sortie []models.BonSortie
	database.DB.Find(&sortie)

	// Formater les dates avant de renvoyer la réponse
	var formattedSortie []map[string]interface{}
	for _, sorti := range sortie {
		formattedSortie = append(formattedSortie, map[string]interface{}{
			"id":           sorti.ID,
			"date":         sorti.Date,
			"materiau_id":  sorti.MateriauID,
			"quantite":     sorti.Quantite,
			"destinataire": sorti.Destinataire,
			"created_at":   sorti.CreatedAt,
			"update_at":    sorti.UpdatedAt,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    formattedSortie,
	})
}
