package main

/*
   Here are the database-related operations.
*/

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// getConn function returns mongodb connection object and context.
func getConn() (*mongo.Client, context.Context) {
	// A new client object is created with the connection string information we have.
	client,
		err := mongo.
		NewClient(options.Client().
			ApplyURI("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false"))
	ctx := context.Background()
	if err != nil {
		log.Printf("Error while opening to new client : %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Error while connecting: %v", err)
	}

	// The client object is pinged to mongoDB.
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Ping yapılamadı: %v", err)
	}
	//Client and context object are returned.
	//Attention should be paid to the use of <defer> in ongoing methods.
	return client, ctx
}

// GetUsers metodu ile crud veritabanındaki userDmos koleksiyonunda yer alan tüm dokümanları çekiyoruz
// Basit bir veri çekme metodu da olsa her ihtimale karşı hata kontrolümüz de var
func GetUsers() ([]*userDmo, error) {
	var userDmos []*userDmo // Döndüreceğimiz array

	client, ctx := getConn()     // MongoDb bağlantı bilgilerini aldık
	defer client.Disconnect(ctx) // Panik olsa da olmasa da metot tamalanırken Disconnect olalım

	db := client.Database("GOCRUD")               //veritabanı nesnesi
	collection := db.Collection("Users")          // koleksiyon nesnesi
	cursor, err := collection.Find(ctx, bson.D{}) // userDmos koleksiyonundaki tüm dokümanları çekmek için kullanılan fonksiyon
	if err != nil {                               // Find metodunun da error dönüşü var o yüzden kontrol etmekte fayda var
		return nil, err
	}
	defer cursor.Close(ctx) // Eğer hata almadan geldiysek sonraki hata durumuna karşın Find ile açılan cursor'u kapattırır

	err = cursor.All(ctx, "es") // Koleksiyonu userDmos'a alıyoruz
	if err != nil {             // All metodunun da error dönüşü var. Kontrol etmek iyi fikir
		log.Println(err)
		return nil, err
	}
	return userDmos, nil // Her şey yolunda gittiyse ;)
}

//AddUser is used to add a new userDmo to the userDmos collection in the <DataBase> database.
func AddUser(u *userDmo) (primitive.ObjectID, error) {
	log.Println("Add User")
	log.Println(u)

	client, ctx := getConn()
	//The method guarantees that the mongodb connection is closed at the end.
	defer client.Disconnect(ctx)
	// The Object Id is created for the Unique ID field and added to the User object.
	u.ID = primitive.NewObjectID()

	// The userDmo record that comes with InsertOne is added to the database.
	// In case of a problem, the err parameter will carry the error information.
	result, err := client.Database("crud").Collection("userDmos").InsertOne(ctx, u)
	//If there is an error, nil is returned with the object ID. The gin-gonic methods will evaluate the error and return the appropriate HTTP message.
	if err != nil {
		log.Printf("An error occurred while adding the record. Hata : %v", err)
		return primitive.NilObjectID, err
	}
	id := result.InsertedID.(primitive.ObjectID)
	// When successful, Object ID information is returned. As there is no error, the second output variable is set to nil
	return id, nil
}

//UpdateUser is used to update a exists user to the user collection in the <DataBase> database.
func UpdateUser(u *userDmo) (int64, error) {
	log.Println("Update User")
	log.Println(u)

	client, ctx := getConn()
	//The method guarantees that the mongodb connection is closed at the end.
	defer client.Disconnect(ctx)

	// The userDmo record that comes with InsertOne is added to the database.
	// In case of a problem, the err parameter will carry the error information.
	result, err := client.
		Database("crud").Collection("userDmos").
		UpdateOne(ctx, u, bson.M{"_id": u.ID})
	//If there is an error, nil is returned with the object ID. The gin-gonic methods will evaluate the error and return the appropriate HTTP message.
	if err != nil {
		log.Printf("An error occurred while adding the record. Hata : %v", err)
		return 0, err
	}

	// When successful, Object ID information is returned. As there is no error, the second output variable is set to nil
	return result.ModifiedCount, nil
}

//DeleteUser is used to delete a exists user to the user collection in the <DataBase> database.
func DeleteUser(uId primitive.ObjectID) (int64, error) {
	log.Println("Delete User")
	log.Println(uId)

	client, ctx := getConn()
	//The method guarantees that the mongodb connection is closed at the end.
	defer client.Disconnect(ctx)

	// The userDmo record that comes with InsertOne is added to the database.
	// In case of a problem, the err parameter will carry the error information.
	result, err := client.
		Database("crud").Collection("userDmos").
		DeleteOne(ctx, bson.M{"_id": uId})
	//If there is an error, nil is returned with the object ID. The gin-gonic methods will evaluate the error and return the appropriate HTTP message.
	if err != nil {
		log.Printf("An error occurred while adding the record. Hata : %v", err)
		return 0, err
	}

	// When successful, Object ID information is returned. As there is no error, the second output variable is set to nil
	return result.DeletedCount, nil
}
