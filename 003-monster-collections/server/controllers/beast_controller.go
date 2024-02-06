package controllers

import (
    "context"
	"log"
    "oo3-monster-collections/server/models"
    "oo3-monster-collections/server/database"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"

)

func GetBeast(c *gin.Context) {
    client, err := database.ConnectToMongoDB()
    if err != nil {
        // Handle error
		log.Fatal(err)
    }
    defer client.Disconnect(context.Background())

    // Access the MongoDB collection
    collection := client.Database("AddYourOwndb").Collection("beasts")

    // Implement your logic to fetch and return beast data
    var result models.Beast
    filter := bson.M{"name": "desired_beast_name"}
    err = collection.FindOne(context.Background(), filter).Decode(&result)
    if err != nil {
        // Handle error
		log.Fatal(err)
    }

    c.JSON(200, result)
}
