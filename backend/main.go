package main

import (
	"context"
	"flag"

	"github.com/Harsh-apk/jwtTest/api"
	"github.com/Harsh-apk/jwtTest/db"
	"github.com/Harsh-apk/jwtTest/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", "localhost:5000", "The Listen Address of the api server")
	flag.Parse()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbUri))
	if err != nil {
		panic(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)
	app.Use(cors.New())

	apiv1 := app.Group("/api/v1", middleware.Auth)
	apiv1.Get("/user", userHandler.HandleGetUser)

	app.Post("/login", userHandler.HandleLogin)
	app.Post("/create", userHandler.HandleCreateAccount)
	app.Static("/", "./public/build")
	app.Listen(*listenAddr)

}
