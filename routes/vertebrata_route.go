package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func VertebrataRoute(router *gin.Engine) {
	router.POST("/vertebrata", controllers.CreateVertebrata())
	router.GET("/vertebrata/:vertebrataId", controllers.GetVertebrata())
	router.PUT("/vertebrata/:vertebrataId", controllers.EditVertebrata())
	router.DELETE("/vertebrata/:vertebrataId", controllers.DeleteVertebrata())
	router.GET("/vertebratas", controllers.GetAllVertebratas())
}
