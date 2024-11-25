package controllers

import (
	"symrise/database"
	"symrise/models"

	"github.com/gofiber/fiber/v2"
)

func AddFournisseur(c *fiber.Ctx) error {
	var data map[string]interface{}

	// Parse le corps de la requête JSON
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erreur de lecture du corps de la requête",
		})
	}

	// Vérification des champs requis
	requiredFields := []string{"nom", "contact", "email", "adresse"}
	for _, field := range requiredFields {
		if _, exists := data[field]; !exists {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Le champ '" + field + "' est requis",
			})
		}
	}

	// Création d'une instance du modèle Fournisseur
	fournisseur := models.Fournisseur{
		Nom:     data["nom"].(string),
		Contact: data["contact"].(string),
		Email:   data["email"].(string),
		Adresse: data["adresse"].(string),
	}

	// Sauvegarde dans la base de données
	if err := database.DB.Create(&fournisseur).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erreur lors de l'ajout du fournisseur",
		})
	}

	// Retourne une réponse de succès
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Fournisseur ajouté avec succès",
		"data":    fournisseur,
	})
}
