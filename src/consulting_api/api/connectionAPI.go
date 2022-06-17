package api

import (
	"consulting_api/configs"
	"consulting_api/controllers"
	"consulting_api/controllers/routes"
	"github.com/gofiber/fiber/v2"
)

func ConnectionAPI(controller *controllers.ConsultingElectoralAuthorityController, electionController *controllers.ConsultingElectoralConfigController) {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Routes.
	// Aqui defino cuales van a ser las rutas accesibles
	routes.PublicRoutesElectoralAuth(app, controller)          // Register a public routes for app.
	routes.PublicRoutesElectionConfig(app, electionController) // Register a public routes for app.
	// Aqui inicializamos el servidor en el puerto 8081
	app.Listen(":8082")
}
