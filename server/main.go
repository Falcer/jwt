package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// User struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Fake table
var users []User

func main() {
	app := fiber.New()

	app.Get("/users", usersHandler)
	app.Post("/login", loginHandler)
	app.Post("/register", registerHandler)

	log.Fatal(app.Listen(":8080"))
}

func usersHandler(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Get All Users",
		"data":    users,
	})
}

func loginHandler(c *fiber.Ctx) error {
	// Login Logic
	u := new(User)

	if err := c.BodyParser(u); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Login Failed",
		})
	}

	for _, user := range users {
		if (*u).Username == user.Username && (*u).Password == user.Password {
			return c.Status(200).JSON(&fiber.Map{
				"success": true,
				"message": "Login successfully",
				"data":    u,
			})
		}
	}

	return c.Status(404).JSON(&fiber.Map{
		"success": false,
		"message": "Username or Password not correct",
	})

}

func registerHandler(c *fiber.Ctx) error {
	// Register Logic
	u := new(User)

	if err := c.BodyParser(u); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Register Failed",
		})
	}

	users = append(users, *u)

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Register",
		"data":    u,
	})
}
