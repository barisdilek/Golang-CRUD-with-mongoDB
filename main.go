package main

// added gin-gonic module
// Also added gin-swagger module for Swagger support
import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "crud/docs" //swag init ile eklenen dokümanın bildirimi
	// Bunu eklememin sebebi çalışma zamanında swagger UI'a gidince aldığım Failed to load spec. hatası

	"github.com/gin-gonic/gin"
)

// @title Golang CRUD with mongoDB Swagger API
// @version 1.0
// @description Service manual
// @termsOfService http://swagger.io/terms/

// @contact.name Barış Dilek
// @contact.email barisdilek@hotmail.com
// @contact.url https://github.com/barisdilek

// @BasePath /api/v1

// @host localhost:5432
func main() {
	router := gin.Default() // instantiated the gin object
	// The route group is determined.
	api := router.Group("/api")
	// group version is determined
	v1 := api.Group("/v1")

	AddEndPoints(v1)

	/*
	   Swagger API is added for documentation support.
	   Swagger UI added at the beginning of endpoint methods will be evaluated by the comment line
	*/
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = router.Run(":5432") // We opened the server from port 5432 to service
}
