package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func AddEndPoints(version *gin.RouterGroup) {
	// Defined a group called User in the version.
	user := version.Group("/Users")
	// Get Http is defined as the Endpoint in the User group.
	user.GET("/", GetAll)

	// To add records to the database with the User object,
	// Post Http is added as an endpoint.
	user.POST("/", Create)

	// Put Http is updated as an endpoint.
	user.PUT("/", Update)

	// Delete Http is deleted as an endpoint.
	user.DELETE("/", Delete)
}

/*
   @Summary : The explanation part for the operation.
   @Produce  : Determined what type of content the HTTP operation works with. example: json.
   @Accept : Determined what type of content the HTTP operation works with. example: json.
   @Param : Indicates that it expects a variable of type userDmo from the body and it is mandatory.
   @Sucess : When the operation is successful, it is stated that HTTP 200 is returned.
   @Failure : HTTP indicates which error type is returned when the operation fails.
   @Router : Indicates that the operation address is triggered by the HTTP POST method.
*/

// GetAll godoc
// @Summary Returns all userDmos
// @Produce json
// @Success 200 {array} userDmo
// @Failure 500
// @Router /Users/ [get]
func GetAll(c *gin.Context) {
	// Data is retrieved from the database.
	var userList, err = GetUsers()
	// The error that occurred during processing to the database is returned.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	// Http 200 : Arrow returns when all goes well.
	c.JSON(http.StatusOK, gin.H{"user": userList})
}

// Create godoc
// @Summary Adds a new userDmo
// @Produce json
// @Accept json
// @Param userDmo
// @Success 200
// @Failure 400
// @Failure 500
// @Router /Users/ [post]
func Create(c *gin.Context) {
	var u userDmo
	// fmt.Println(u
	// Checking the json body content to match the userDmo object.
	if err := c.ShouldBindJSON(&u); err != nil {
		log.Print(err)
		// When there is a problem with the JSON content, HTTP 400 is returned with the error message.
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	// The content in the resulting userDmo object is added to the database.
	id, err := AddUser(&u)
	// The error that occurred during processing to the database is returned.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	u.ID = id
	// Http 200 : Arrow returns when all goes well.
	c.JSON(http.StatusOK, gin.H{"added": u})
}

// Update godoc
// @Summary Update a userDmo
// @Produce json
// @Accept json
// @Param userDmo body userDmo true "User infos"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /Users/ [Put]
func Update(c *gin.Context) {
	var u userDmo
	// Checking the json body content to match the userDmo object.
	if err := c.ShouldBindJSON(&u); err != nil {
		log.Print(err)
		// When there is a problem with the JSON content, HTTP 400 is returned with the error message.
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	// The content in the resulting userDmo object is added to the database.
	count, err := UpdateUser(&u)
	// The error that occurred during processing to the database is returned.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	// Http 200 : Arrow returns when all goes well.
	c.JSON(http.StatusOK, gin.H{"added": count})
}

// Delete godoc
// @Summary Delete a userDmo
// @Produce json
// @Accept json
// @Param  userDmo body userDmo true "User Id"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /Users/ [Delete]
func Delete(c *gin.Context) {
	var uId primitive.ObjectID
	// Checking the json body content to match the userDmo object.
	if err := c.ShouldBindJSON(&uId); err != nil {
		log.Print(err)
		// When there is a problem with the JSON content, HTTP 400 is returned with the error message.
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	// The content in the resulting userDmo object is added to the database.
	count, err := DeleteUser(uId)
	// The error that occurred during processing to the database is returned.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	// Http 200 : Arrow returns when all goes well.
	c.JSON(http.StatusOK, gin.H{"deleted": count})
}

/*
   Specifies that {array} will return a list of userDmos.
*/
