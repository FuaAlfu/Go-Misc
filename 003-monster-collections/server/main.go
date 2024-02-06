package main

import (
    "oo3-monster-collections/server/controllers"

    "github.com/gin-gonic/gin"

)

func main() {
    r := gin.Default()

    // Define routes
    r.GET("/beast", controllers.GetBeast)

    // Run the server
    r.Run(":8080")
}
