package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
   The data model object in the Go language of the userDmo object in mongoDB.
   The ID can also be expressed with one of bson's primitive types for mongodb.
*/
type userDmo struct {
	ID      primitive.ObjectID
	Name    string
	SurName string
	Pass    string
	Email   string
}
