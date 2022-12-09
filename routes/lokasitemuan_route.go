package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func LokasiTemuanRoute(router *gin.Engine) {
	router.POST("/lokasitemuan", controllers.CreateLokasiTemuan())
	router.GET("/lokasitemuan/:lokasitemuanId", controllers.GetLokasiTemuan())
	router.PUT("/lokasitemuan/:lokasitemuanId", controllers.EditLokasiTemuan())
	router.DELETE("/lokasitemuan/:lokasitemuanId", controllers.DeleteLokasiTemuan())
	router.GET("/lokasitemuans", controllers.GetAllLokasiTemuans())
}

func KoordinatRoute(router *gin.Engine) {
	router.GET("/koordinat/:longitude/:latitude", controllers.GetKoordinat())
}
