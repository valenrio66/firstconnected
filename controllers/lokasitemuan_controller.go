package controllers

import (
	"context"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var lokasitemuanCollection *mongo.Collection = configs.GetCollection(configs.DB, "lokasitemuan")
var validate_lokasitemuan = validator.New()

func CreateLokasiTemuan() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var lokasitemuan models.LokasiTemuan
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&lokasitemuan); err != nil {
			c.JSON(http.StatusBadRequest, responses.LokasiTemuanResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_lokasitemuan.Struct(&lokasitemuan); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LokasiTemuanResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newLokasiTemuan := models.LokasiTemuan{
			Lokasi:    lokasitemuan.Lokasi,
			Kelurahan: lokasitemuan.Kelurahan,
			Kecamatan: lokasitemuan.Kecamatan,
			Kota:      lokasitemuan.Kota,
			Provinsi:  lokasitemuan.Provinsi,
		}

		result, err := lokasitemuanCollection.InsertOne(ctx, newLokasiTemuan)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LokasiTemuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.LokasiTemuanResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetLokasiTemuan() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		lokasitemuanId := c.Param("lokasitemuanId")
		var lokasitemuan models.LokasiTemuan
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(lokasitemuanId)

		err := lokasitemuanCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&lokasitemuan)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LokasiTemuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.LokasiTemuanResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": lokasitemuan}})
	}
}

func EditLokasiTemuan() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		lokasitemuanId := c.Param("lokasitemuanId")
		var lokasitemuan models.LokasiTemuan
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(lokasitemuanId)

		//validate the request body
		if err := c.BindJSON(&lokasitemuan); err != nil {
			c.JSON(http.StatusBadRequest, responses.LokasiTemuanResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_lokasitemuan.Struct(&lokasitemuan); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LokasiTemuanResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"lokasi": lokasitemuan.Lokasi, "kelurahan": lokasitemuan.Kelurahan, "kecamatan": lokasitemuan.Kecamatan, "kota": lokasitemuan.Kota, "provinsi": lokasitemuan.Provinsi}
		result, err := lokasitemuanCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LokasiTemuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated LokasiTemuan details
		var updatedLokasiTemuan models.LokasiTemuan
		if result.MatchedCount == 1 {
			err := lokasitemuanCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedLokasiTemuan)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LokasiTemuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.LokasiTemuanResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedLokasiTemuan}})
	}
}

func DeleteLokasiTemuan() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		lokasitemuanId := c.Param("lokasitemuanId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(lokasitemuanId)

		result, err := lokasitemuanCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LokasiTemuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.LokasiTemuanResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "LokasiTemuan with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.LokasiTemuanResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "LokasiTemuan successfully deleted!"}},
		)
	}
}

func GetAllLokasiTemuans() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var lokasitemuans []models.LokasiTemuan
		defer cancel()

		results, err := lokasitemuanCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LokasiTemuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleLokasiTemuan models.LokasiTemuan
			if err = results.Decode(&singleLokasiTemuan); err != nil {
				c.JSON(http.StatusInternalServerError, responses.LokasiTemuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			lokasitemuans = append(lokasitemuans, singleLokasiTemuan)
		}

		c.JSON(http.StatusOK,
			responses.LokasiTemuanResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": lokasitemuans}},
		)
	}
}
