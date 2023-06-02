package main

import (
	"log"

	"github.com/Ryzhon/go_chat/pkg/commonn/config"
	"github.com/Ryzhon/go_chat/pkg/commonn/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed at config", err)
	}
	app := fiber.New()
	db := db.Init(c.DBUrl)

	books.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
