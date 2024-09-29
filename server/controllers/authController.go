package controllers

import (
	"context"
	"log"
	"server-backend/config"
	"server-backend/models"
	"server-backend/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = config.ConnectDB().Database("test").Collection("users")

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if user email already exists in the database
	var existingUser models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)

	// If a user with the same email is found, return without creating a new user
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "User with this email already exists"})
	}

	// Hash Password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	log.Println("User pre-insertion > ", user)
	log.Println("UserCollection pre-insertion > ", userCollection)

	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to register user"})
	}

	log.Println("result post-insertion > ", result) // This must contains id_ & email

	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to generate token"})
	}

	user.Token = token
	return c.Status(fiber.StatusOK).JSON(result)
}

func Login(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Check Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to generate token"})
	}

	user.Token = token
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

func Profile(c *fiber.Ctx) error {
	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"message": "Access granted to the protected route!",
	// })

	userEmail := c.Locals("user_email").(string)
	CLAIMs := c.Locals("CLAIMs")

	log.Println("userEmail", userEmail)
	log.Println("CLAIMs", CLAIMs)

	// Find the user by email
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching user data"})
	}

	// Return user details (without unhashed password for security reasons)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})

}
