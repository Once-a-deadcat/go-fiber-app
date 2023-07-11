package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func middleware(c *fiber.Ctx) error {
	fmt.Println("Don't mind me!")
	return c.Next()
}

func handler(c *fiber.Ctx) error {
	return c.SendString(c.Path())
}

func main() {
	app := fiber.New()

	// GET request
	app.Get("/api/user", func(c *fiber.Ctx) error {
		user := User{"John Doe", "john.doe@example.com"}
		return c.JSON(user)
	})

	// POST request
	app.Post("/api/user", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			return c.Status(503).SendString(err.Error())
		}

		fmt.Printf("Posted User: %+v\n", user)
		return c.JSON(user)
	})

	// GET /api/register
	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})

	// GET /dictionary.txt
	app.Get("/:file.:ext", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => 📃 dictionary.txt
	})

	// GET /john/75
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old. gender: %s", c.Params("name"), c.Params("age"), c.Query("gender"))
		return c.SendString(msg) // => 👴 john is 75 years old
	})

	// Root API route
	api := app.Group("/api", middleware) // /api

	// API v1 routes
	v1 := api.Group("/v1", middleware) // /api/v1
	v1.Get("/list", handler)           // /api/v1/list
	v1.Get("/user", handler)           // /api/v1/user

	// API v2 routes
	v2 := api.Group("/v2", middleware) // /api/v2
	v2.Get("/list", handler)           // /api/v2/list
	v2.Get("/user", handler)           // /api/v2/user

	app.Listen(":3000")
}
