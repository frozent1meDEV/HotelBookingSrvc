package main

import (
	"HotelBookingSrvc/api"
	"HotelBookingSrvc/db"
	_ "HotelBookingSrvc/types"
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {

	listenAddr := flag.String("ListenAddr", ":5000", "listen address")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	apiv1.Post("/User", userHandler.HandlePostUser)
	apiv1.Get("/User", userHandler.HandleGetUsers)
	apiv1.Get("/User/:id", userHandler.HandleGetUser)
	app.Listen(*listenAddr)
}
