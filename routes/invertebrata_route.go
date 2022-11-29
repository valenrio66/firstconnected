package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func InvertebrataRoute(router *gin.Engine) {
	router.POST("/invertebrata", controllers.CreateInvertebrata())
	router.GET("/invertebrata/:invertebrataId", controllers.GetInvertebrata())
	router.PUT("/invertebrata/:invertebrataId", controllers.EditInvertebrata())
	router.DELETE("/invertebrata/:invertebrataId", controllers.DeleteInvertebrata())
	router.GET("/invertebratas", controllers.GetAllInvertebratas())
}
