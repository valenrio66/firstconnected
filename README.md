# Golang Gin MongoDB Rest API

Create Rest API Using Go Language with Gin Framework and MongoDB.

## Getting Started
1. Create Folder for Starting the Project
2. Initialize Go Module to Manage Project Dependencies by Running The Command Below
```sh
go mod init your-folder-name
```
3. Install required dependencies now (first time)
```sh
go get -u github.com/gin-gonic/gin 
go get go.mongodb.org/mongo-driver/mongo 
go get github.com/joho/godotenv 
go get github.com/go-playground/validator/v10
```
4. After the project dependencies installed, create `main.go` file in the root directory and add the snippet below
```go
package main

import "github.com/gin-gonic/gin"

func main() {
        router := gin.Default()

        router.GET("/", func(c *gin.Context) {
                c.JSON(200, gin.H{
                        "data": "Hello from Gin-gonic & mongoDB",
                })
        })

        router.Run("localhost:8080") 
}
```
5. Run app
```sh
go run main.go
```

6. App is running on port: 8080
7. Test it using Thunder Client in VSCode or Using Postman with the URL is
```sh
http://localhost:8080
```