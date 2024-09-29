package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"os"
	"server-backend/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	} else {
		log.Println("Successful loading .env file")
	}
}

func main() {

	fmt.Println("Hello from main")

	app := fiber.New()

	routes.AuthRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
