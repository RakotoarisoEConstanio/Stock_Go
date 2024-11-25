package controllers

import (
	"fmt"
	"symrise/database"
	"symrise/models"

	"github.com/gofiber/fiber/v2"
)

func GetMateriaux(c *fiber.Ctx) error {
	var materiaux []models.Materiaux
	database.DB.Find(&materiaux)

	// Formater les dates avant de renvoyer la réponse
	var formattedMateriaux []map[string]interface{}
	for _, materiau := range materiaux {
		formattedMateriaux = append(formattedMateriaux, map[string]interface{}{
			"id":            materiau.ID,
			"nom":           materiau.Nom,
			"description":   materiau.Description,
			"stock_initial": materiau.StockInitial,
			"stock_actuel":  materiau.StockActuel,
			"seuil_min":     materiau.SeuilMin,
			"created_at":    materiau.CreatedAt,
			"update_at":     materiau.UpdatedAt,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    formattedMateriaux,
	})
}

func AddMateriaux(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erreur de lecture du corps de la requête ",
		})
	}

	requiredFields := []string{"nom", "description", "stock_initial", "stock_actuel", "seuil_min"}
	for _, field := range requiredFields {
		if _, exists := data[field]; !exists {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("Le champ '%s' est requis", field),
			})
		}
	}

	nom, ok := data["nom"].(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Le champ 'nom' est requis et doit être une chaîne de caractères",
		})
	}

	// Vérification et conversion du champ "description" (string)
	description, ok := data["description"].(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Le champ 'description' est requis et doit être une chaîne de caractères",
		})
	}

	stockInitial, ok := data["stock_initial"].(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Valeur invalide pour le Stock Initial",
		})
	}

	stockActuel, ok := data["stock_actuel"].(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Valeur invalide pour le Stock Actuel",
		})
	}

	Seuilmin, ok := data["seuil_min"].(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Valeur invalide pour le Seuil minimal",
		})
	}

	materiaux := models.Materiaux{
		Nom:          nom,
		Description:  description,
		StockInitial: stockInitial,
		StockActuel:  stockActuel,
		SeuilMin:     Seuilmin,
	}

	database.DB.Create(&materiaux)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Materiaux ajoutée avec succées",
		"data":    materiaux,
	})
}
