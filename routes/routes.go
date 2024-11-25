package routes

import (
	"symrise/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/logout", controllers.Logout)

	//  matériaux

	app.Post("/api/addMateriaux", controllers.AddMateriaux)
	app.Get("/api/getMateriaux", controllers.GetMateriaux)

	// Bon d'entrée

	app.Post("/api/addbonentree", controllers.AddEntree)

	// Fournisseur

	app.Post("/api/addfournisseur", controllers.AddFournisseur)

	// bon de sortie
	app.Get("/api/getbonsortie", controllers.GetBonsortie)
	app.Post("/api/addbonsortie", controllers.AddSortie)
}
