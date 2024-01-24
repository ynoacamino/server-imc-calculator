package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ynoacamino/server-imc-calculator/db"
	"github.com/ynoacamino/server-imc-calculator/structs"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Post("/add", func(c *fiber.Ctx) error {
		user := new(structs.User)

		err := c.BodyParser(user)
		if err != nil {
			return err
		}

		if parserUser(user) {
			connectDB := db.GetDB()
			_, err = connectDB.Exec("INSERT INTO imc (peso, talla, response) VALUES(?, ?, ?)", user.Peso, user.Talla, user.Response)

			if err != nil {
				return err
			}

			return c.JSON(fiber.Map{
				"response": "ok",
			})
		}

		return c.JSON(fiber.Map{
			"response": "no",
		})
	})

	listenError := app.Listen(":" + os.Getenv("PORT"))
	if listenError != nil {
		log.Fatal("Connect error")
	}
}

func parserUser(user *structs.User) bool {
	if (user.Peso > 5 && user.Peso < 300) && (user.Talla > 50 && user.Talla < 300) {
		return true
	}
	return false
}
