package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/RavshanovUsmonbek/hotel-reservation/types"
	"strconv"
)

// GetUserByIdHandler godoc
// @Summary Get a user by ID
// @Description Get user details by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} types.User
// @Failure 400 {object} map[string]string
// @Router /users/{id} [get]
func GetUserByIdHandler(c *fiber.Ctx) error {
	userId := c.Params("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	user := types.User{
		ID:       id,
		Username: "testuser",
		Email:    "TesEmail",
	}
	return c.JSON(user)
}

// GetAllUsersHandler godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} types.User
// @Router /users [get]
func GetAllUsersHandler(c *fiber.Ctx) error {
	users := []types.User{
		{ID: 1, Username: "user1", Email: "TestEmail1"},
		{ID: 2, Username: "user2", Email: "TestEmail2"},
		{ID: 3, Username: "user3", Email: "TestEmail3"},
	}
	return c.JSON(users)
}

// CreateUserHandler godoc
// @Summary Create a new user
// @Description Create a user from JSON body
// @Tags users
// @Accept json
// @Produce json
// @Param user body types.User true "User object"
// @Success 201 {object} types.User
// @Failure 400 {object} map[string]string
// @Router /users [post]
func CreateUserHandler(c *fiber.Ctx) error {
	var user types.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	user.ID = 4
	return c.Status(201).JSON(user)
}

// UpdateUserHandler godoc
// @Summary Update an existing user
// @Description Update a user by ID with JSON body
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body types.User true "User object"
// @Success 200 {object} types.User
// @Failure 400 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUserHandler(c *fiber.Ctx) error {
	var user types.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	userId := c.Params("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	user.ID = id
	return c.JSON(user)
}

// DeleteUserHandler godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [delete]
func DeleteUserHandler(c *fiber.Ctx) error {
	userId := c.Params("id")
	return c.JSON(fiber.Map{"message": "User deleted successfully", "userId": userId})
}
