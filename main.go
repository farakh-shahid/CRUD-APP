package main

import (
	"context"
	"log"
	"os"

	"github.com/farakh-shahid/CRUD-APP/controllers"
	"github.com/farakh-shahid/CRUD-APP/routes"
	"github.com/farakh-shahid/CRUD-APP/services"
	"github.com/farakh-shahid/CRUD-APP/utils/constants"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx := context.TODO()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(constants.ErrorLoadingEnv)
	}

	mongoURI := os.Getenv("MONGO_URI")

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(constants.ErrorWhileConnectingWithMongoDB, err)
	}
	defer mongoClient.Disconnect(ctx)

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(constants.ErrorWhilePingToMongoDB, err)
	}

	userCollection := mongoClient.Database("userdb").Collection("users")
	userService := services.NewUserService(userCollection, ctx)
	userController := &controllers.UserController{
		UserService: userService,
	}

	router := gin.Default()
	routes.RegisterUserRoutes(router.Group("/api/v1"), userController)
	PORT := os.Getenv("PORT")
	err = router.Run(":" + PORT)
	if err != nil {
		log.Fatal(constants.ErrorWhileStartingServer, err)
	}
}
