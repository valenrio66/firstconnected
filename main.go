package main

import (

	//add this
	"context"
	"log"
	"net/http"

	"gin-mongo-api/configs"
	"gin-mongo-api/controllers"
	"gin-mongo-api/routes"
	"gin-mongo-api/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client

	userService         services.UserService
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	authCollection      *mongo.Collection
	authService         services.AuthService
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController
)

func main() {
	configs, err := configs.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer mongoclient.Disconnect(ctx)

	authCollection = mongoclient.Database("dbmuseum").Collection("user")
	userService = services.NewUserServiceImpl(authCollection, ctx)
	authService = services.NewAuthService(authCollection, ctx)
	AuthController = controllers.NewAuthController(authService, userService)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(userService)
	UserRouteController = routes.NewRouteUserController(UserController)

	server = gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"https://localhost:8888", "http://localhost:3000"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	AuthRouteController.AuthRoute(router, userService)
	UserRouteController.UserRoute(router, userService)
	log.Fatal(server.Run(":" + configs.Port))
	// router := gin.Default()

	// //run database
	// configs.ConnectDB()

	// //routes
	// // routes.UserRoute(router)              //add this
	// routes.InvertebrataRoute(router)      //add this
	// routes.VertebrataRoute(router)        //add this
	// routes.FosilRoute(router)             //add this
	// routes.BatuanRoute(router)            //add this
	// routes.SumberDayaGeologiRoute(router) //add this
	// routes.LokasiTemuanRoute(router)      //add this
	// routes.KoordinatRoute(router)         //add this

	// router.Run(":" + SetPort())
}

// func SetPort() string {
// 	port := os.Getenv("PORT")
// 	if len(port) == 0 {
// 		port = "80"
// 	}
// 	return port
// }
