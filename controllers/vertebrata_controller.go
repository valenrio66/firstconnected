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

var vertebrataCollection *mongo.Collection = configs.GetCollection(configs.DB, "vertebrata")
var validate_vertebrata = validator.New()

func CreateVertebrata() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var vertebrata models.Vertebrata
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&vertebrata); err != nil {
			c.JSON(http.StatusBadRequest, responses.VertebrataResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_vertebrata.Struct(&vertebrata); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.VertebrataResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newVertebrata := models.Vertebrata{
			Id:               primitive.NewObjectID(),
			Nama:             vertebrata.Nama,
			Lokasi_Ditemukan: vertebrata.Lokasi_Ditemukan,
			Waktu_Ditemukan:  vertebrata.Waktu_Ditemukan,
		}

		result, err := vertebrataCollection.InsertOne(ctx, newVertebrata)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.VertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.VertebrataResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetVertebrata() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		vertebrataId := c.Param("vertebrataId")
		var vertebrata models.Vertebrata
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(vertebrataId)

		err := vertebrataCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&vertebrata)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.VertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.VertebrataResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": vertebrata}})
	}
}

func EditVertebrata() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		vertebrataId := c.Param("vertebrataId")
		var vertebrata models.Vertebrata
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(vertebrataId)

		//validate the request body
		if err := c.BindJSON(&vertebrata); err != nil {
			c.JSON(http.StatusBadRequest, responses.VertebrataResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_vertebrata.Struct(&vertebrata); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.VertebrataResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"nama": vertebrata.Nama, "lokasi_ditemukan": vertebrata.Lokasi_Ditemukan, "waktu_ditemukan": vertebrata.Waktu_Ditemukan}
		result, err := vertebrataCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.VertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated Vertebrata details
		var updatedVertebrata models.Vertebrata
		if result.MatchedCount == 1 {
			err := vertebrataCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedVertebrata)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.VertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.VertebrataResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedVertebrata}})
	}
}

func DeleteVertebrata() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		vertebrataId := c.Param("vertebrataId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(vertebrataId)

		result, err := vertebrataCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.VertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.VertebrataResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Vertebrata with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.VertebrataResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Vertebrata successfully deleted!"}},
		)
	}
}

func GetAllVertebratas() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var vertebratas []models.Vertebrata
		defer cancel()

		results, err := vertebrataCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.VertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleVertebrata models.Vertebrata
			if err = results.Decode(&singleVertebrata); err != nil {
				c.JSON(http.StatusInternalServerError, responses.VertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			vertebratas = append(vertebratas, singleVertebrata)
		}

		c.JSON(http.StatusOK,
			responses.VertebrataResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": vertebratas}},
		)
	}
}
