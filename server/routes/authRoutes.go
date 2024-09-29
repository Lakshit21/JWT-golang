package routes

import (
	"server-backend/controllers"
	"server-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	// Protected route
	app.Get("/profile", middleware.JWTProtected(), controllers.Profile)
}
