package main

import (
	"api/Battery-Tracking-go/controller"
	"api/Battery-Tracking-go/service"
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
)

var (
 	server      *gin.Engine
 	us          service.RelatedNumberService
 	uc          controller.RelatedNumberController
 	ctx         context.Context
	userc       *mongo.Collection
 	mongoclient *mongo.Client
 	err         error
)

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://username:password@202.151.177.207:27017").SetAuth(options.Credential{
		AuthSource: "battery",
		Username:   "dev",
		Password:   "12345678",
	})
	// Auto Migrate the struct
	
	

	// Connect to MongoDB server.
	mongoclient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection.
	err = mongoclient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB server.")


	userc = mongoclient.Database("battery").Collection("Battert-Tracking")
	us = service.NewRelatedNumberService(userc, ctx)
	uc = controller.New(us)
	server = gin.Default()
}


func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	uc.RegisterUserRouts(basepath)

	log.Fatal(server.Run(":9098"))
}

