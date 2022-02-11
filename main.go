
package main

import (
		"go-service-101/configs"
		"go-service-101/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    //run database
    configs.ConnectDB()
		routes.UserRoute(router)
    router.Run("localhost:8080")
}