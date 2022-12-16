package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes" //add this
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.InvertebrataRoute(router)      //add this
	routes.VertebrataRoute(router)        //add this
	routes.FosilRoute(router)             //add this
	routes.BatuanRoute(router)            //add this
	routes.SumberDayaGeologiRoute(router) //add this
	routes.LokasiTemuanRoute(router)      //add this
	routes.KoordinatRoute(router)         //add this

	router.Run(":" + SetPort())
}

func SetPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}
	return port
}
