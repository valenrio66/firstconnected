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

var sumberdayageologiCollection *mongo.Collection = configs.GetCollection(configs.DB, "sumberdayageologi")
var validate_sumberdayageologi = validator.New()

func CreateSumberDayaGeologi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var sumberdayageologi models.SumberDayaGeologi
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&sumberdayageologi); err != nil {
			c.JSON(http.StatusBadRequest, responses.SumberDayaGeologiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_sumberdayageologi.Struct(&sumberdayageologi); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.SumberDayaGeologiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newSumberDayaGeologi := models.SumberDayaGeologi{
			Id:               primitive.NewObjectID(),
			NoRegister:       sumberdayageologi.NoRegister,
			NoInventaris:     sumberdayageologi.NoInventaris,
			KodeBmn:          sumberdayageologi.KodeBmn,
			NupBmn:           sumberdayageologi.NupBmn,
			MerkBmn:          sumberdayageologi.MerkBmn,
			KelompokKoleksi:  sumberdayageologi.KelompokKoleksi,
			JenisKoleksi:     sumberdayageologi.JenisKoleksi,
			SubJenisKoleksi:  sumberdayageologi.SubJenisKoleksi,
			KodeJenisKoleksi: sumberdayageologi.KodeJenisKoleksi,
			RuangSimpan:      sumberdayageologi.RuangSimpan,
			LokasiSimpan:     sumberdayageologi.LokasiSimpan,
			Kondisi:          sumberdayageologi.Kondisi,
			NamaKoleksi:      sumberdayageologi.NamaKoleksi,
			Keterangan:       sumberdayageologi.Keterangan,
			LokasiTemuan:     sumberdayageologi.LokasiTemuan,
			Pulau:            sumberdayageologi.Pulau,
			CaraPerolehan:    sumberdayageologi.CaraPerolehan,
			TahunPerolehan:   sumberdayageologi.TahunPerolehan,
			Kolektor:         sumberdayageologi.Kolektor,
			Kepemilikan:      sumberdayageologi.Kepemilikan,
			Operator:         sumberdayageologi.Operator,
			TanggalDicatat:   sumberdayageologi.TanggalDicatat,
			NilaiPerolehan:   sumberdayageologi.NilaiPerolehan,
			NilaiBuku:        sumberdayageologi.NilaiBuku,
		}

		result, err := sumberdayageologiCollection.InsertOne(ctx, newSumberDayaGeologi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.SumberDayaGeologiResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetSumberDayaGeologi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		sumberdayageologiId := c.Param("sumberdayageologiId")
		var sumberdayageologi models.SumberDayaGeologi
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(sumberdayageologiId)

		err := sumberdayageologiCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&sumberdayageologi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.SumberDayaGeologiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": sumberdayageologi}})
	}
}

func EditSumberDayaGeologi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		sumberdayageologiId := c.Param("sumberdayageologiId")
		var sumberdayageologi models.SumberDayaGeologi
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(sumberdayageologiId)

		//validate the request body
		if err := c.BindJSON(&sumberdayageologi); err != nil {
			c.JSON(http.StatusBadRequest, responses.SumberDayaGeologiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_sumberdayageologi.Struct(&sumberdayageologi); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.SumberDayaGeologiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"no_register": sumberdayageologi.NoRegister, "no_inventaris": sumberdayageologi.NoInventaris, "kode_bmn": sumberdayageologi.KodeBmn, "nup_bmn": sumberdayageologi.NupBmn, "merk_bmn": sumberdayageologi.MerkBmn, "kelompok_koleksi": sumberdayageologi.KelompokKoleksi, "jenis_koleksi": sumberdayageologi.JenisKoleksi, "sub_jenis_koleksi": sumberdayageologi.SubJenisKoleksi, "kode_jenis_koleksi": sumberdayageologi.KodeJenisKoleksi, "ruang_simpan": sumberdayageologi.RuangSimpan, "lokasi_simpan": sumberdayageologi.LokasiSimpan, "kondisi": sumberdayageologi.Kondisi, "nama_koleksi": sumberdayageologi.NamaKoleksi, "keterangan": sumberdayageologi.Keterangan, "lokasi_temuan": sumberdayageologi.LokasiTemuan, "pulau": sumberdayageologi.Pulau, "cara_perolehan": sumberdayageologi.CaraPerolehan, "tahun_perolehan": sumberdayageologi.TahunPerolehan, "kolektor": sumberdayageologi.Kolektor, "kepemilikan": sumberdayageologi.Kepemilikan, "operator": sumberdayageologi.Operator, "tanggal_dicatat": sumberdayageologi.TanggalDicatat, "nilai_perolehan": sumberdayageologi.NilaiPerolehan, "nilai_buku": sumberdayageologi.NilaiBuku}
		result, err := sumberdayageologiCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated SumberDayaGeologi details
		var updatedSumberDayaGeologi models.SumberDayaGeologi
		if result.MatchedCount == 1 {
			err := sumberdayageologiCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedSumberDayaGeologi)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.SumberDayaGeologiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedSumberDayaGeologi}})
	}
}

func DeleteSumberDayaGeologi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		sumberdayageologiId := c.Param("sumberdayageologiId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(sumberdayageologiId)

		result, err := sumberdayageologiCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.SumberDayaGeologiResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "SumberDayaGeologi with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.SumberDayaGeologiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "SumberDayaGeologi successfully deleted!"}},
		)
	}
}

func GetAllSumberDayaGeologis() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var sumberdayageologis []models.SumberDayaGeologi
		defer cancel()

		results, err := sumberdayageologiCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleSumberDayaGeologi models.SumberDayaGeologi
			if err = results.Decode(&singleSumberDayaGeologi); err != nil {
				c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			sumberdayageologis = append(sumberdayageologis, singleSumberDayaGeologi)
		}

		c.JSON(http.StatusOK,
			responses.SumberDayaGeologiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": sumberdayageologis}},
		)
	}
}
