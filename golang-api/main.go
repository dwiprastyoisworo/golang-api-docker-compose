package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang-api/app/dto"
	"golang-api/config"
	migrationPostgres "golang-api/migration"
	"log"
)

func main() {
	// setup fiber
	app := fiber.New()

	// get postgres database connection
	postgres := config.NewPostgresConnection(config.GlobalEnv.Postgres)
	db := postgres.Connect()

	//run migration process
	migration := migrationPostgres.NewMigration(config.GlobalEnv.Postgres)
	migration.Run()

	// sample api get
	app.Get("/", func(c *fiber.Ctx) error {
		var users []dto.Users
		err := db.Find(&users).Error
		if err != nil {
			log.Print(err)
			return c.SendStatus(fiber.StatusBadGateway)
		}
		return c.Status(fiber.StatusOK).JSON(users)
	})

	// set port listen
	app.Listen(fmt.Sprintf(":%s", config.GlobalEnv.Port))
}
