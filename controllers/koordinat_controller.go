package controllers

import (
	"context"
	"fmt"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var villagesCollection *mongo.Collection = configs.GetCollection(configs.DB, "villages")

func GetKoordinat() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var villages models.Villages
		longitude := c.Param("longitude")
		long, _ := strconv.ParseFloat(longitude, 64)
		latitude := c.Param("latitude")
		lat, _ := strconv.ParseFloat(latitude, 64)
		defer cancel()

		filter := bson.M{
			"border": bson.M{
				"$geoIntersects": bson.M{
					"$geometry": bson.M{
						"type":        "Point",
						"coordinates": []float64{long, lat},
					},
				},
			},
		}
		fmt.Print("Ini inputan long", long)
		fmt.Print("Ini inputan lat", lat)
		err := villagesCollection.FindOne(ctx, filter).Decode(&villages)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.KoordinatResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.KoordinatResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": villages}})
	}
}
