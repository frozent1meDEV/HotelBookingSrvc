package main

import (
	"HotelBookingSrvc/api"
	"HotelBookingSrvc/db"
	"HotelBookingSrvc/types"
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {

	listenAddr := flag.String("ListenAddr", ":5000", "listen address")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/Users", userHandler.HandleGetUsers)
	apiv1.Get("/Users/:id", userHandler.HandleGetUser)
	app.Listen(*listenAddr)
}
