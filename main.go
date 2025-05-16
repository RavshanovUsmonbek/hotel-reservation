// filepath: /Users/usmonbek/Programming/Learning/Go/hotel-reservation/main.go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/swagger" // Swagger handler for Fiber
    _ "github.com/RavshanovUsmonbek/hotel-reservation/docs" // Swagger docs
    "github.com/RavshanovUsmonbek/hotel-reservation/api"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("This is a cool app!")
    })

    // Swagger route
    app.Get("/swagger/*", swagger.HandlerDefault) // Serve Swagger UI

    app.Get("/users/:id", api.GetUserByIdHandler)
    app.Get("/users", api.GetAllUsersHandler)
    app.Post("/users", api.CreateUserHandler)
    app.Put("/users/:id", api.UpdateUserHandler)
    app.Delete("/users/:id", api.DeleteUserHandler)

    app.Listen(":3000")
}