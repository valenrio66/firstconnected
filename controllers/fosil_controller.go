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

var fosilCollection *mongo.Collection = configs.GetCollection(configs.DB, "fosil")
var validate_fosil = validator.New()

func CreateFosil() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var fosil models.Fosil
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&fosil); err != nil {
			c.JSON(http.StatusBadRequest, responses.FosilResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_fosil.Struct(&fosil); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.FosilResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newFosil := models.Fosil{
			Id:             primitive.NewObjectID(),
			NoRegister:     fosil.NoRegister,
			NoInventaris:   fosil.NoInventaris,
			NamaKoleksi:    fosil.NamaKoleksi,
			LokasiTemuan:   fosil.LokasiTemuan,
			TahunPerolehan: fosil.TahunPerolehan,
			Determinator:   fosil.Determinator,
			Keterangan:     fosil.Keterangan,
		}

		result, err := fosilCollection.InsertOne(ctx, newFosil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.FosilResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetFosil() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		fosilId := c.Param("fosilId")
		var fosil models.Fosil
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(fosilId)

		err := fosilCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&fosil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.FosilResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": fosil}})
	}
}

func EditFosil() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		fosilId := c.Param("fosilId")
		var fosil models.Fosil
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(fosilId)

		//validate the request body
		if err := c.BindJSON(&fosil); err != nil {
			c.JSON(http.StatusBadRequest, responses.FosilResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_fosil.Struct(&fosil); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.FosilResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"No Register": fosil.NoRegister, "No Inventaris": fosil.NoInventaris, "Nama Koleksi": fosil.NamaKoleksi, "Lokasi Temuan": fosil.LokasiTemuan, "Tahun Perolehan": fosil.TahunPerolehan, "determinator": fosil.Determinator, "keterangan": fosil.Keterangan}
		result, err := fosilCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated Fosil details
		var updatedFosil models.Fosil
		if result.MatchedCount == 1 {
			err := fosilCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedFosil)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.FosilResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedFosil}})
	}
}

func DeleteFosil() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		fosilId := c.Param("fosilId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(fosilId)

		result, err := fosilCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.FosilResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Fosil with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.FosilResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Fosil successfully deleted!"}},
		)
	}
}

func GetAllFosils() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var fosils []models.Fosil
		defer cancel()

		results, err := fosilCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleFosil models.Fosil
			if err = results.Decode(&singleFosil); err != nil {
				c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			fosils = append(fosils, singleFosil)
		}

		c.JSON(http.StatusOK,
			responses.FosilResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": fosils}},
		)
	}
}
