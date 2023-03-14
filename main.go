package main

import (
    "github.com/jrschmidtt/go_mongo/configs" //add this
	"github.com/jrschmidtt/go_mongo/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    //run database
    configs.ConnectDB()

	//run routes
	routes.UserRoute(router)

    router.Run("localhost:6000")
}