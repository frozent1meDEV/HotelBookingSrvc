package main

import (
	"HotelBookingSrvc/api"
	"flag"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("ListenAddr", ":5000", "listen address")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/Users", api.HandleGetUsers)
	apiv1.Get("/Users/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}

func handleHotel(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working at port  :5000"})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Jhon Doe"})
}
