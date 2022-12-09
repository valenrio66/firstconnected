package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes" //add this

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)         //add this
	routes.InvertebrataRoute(router) //add this
	routes.VertebrataRoute(router)   //add this
	routes.FosilRoute(router)        //add this
	routes.LokasiTemuanRoute(router)
	routes.KoordinatRoute(router)

	router.Run("localhost:8080")
}
