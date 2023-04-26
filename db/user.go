package db

import (
	"fmt"
	"go-practice/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers() ([]models.User, error) {

	ConnectDB()

	coll := Client.Database("create-app").Collection("users")
	var users []models.User
	cursor, err := coll.Find(Ctx, bson.D{})
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found")
	}

	if err != nil {
		fmt.Println("error getting document: ", err)
	}

	if err = cursor.All(Ctx, &users); err != nil {
		fmt.Println("error parsing:", err)
	}

	fmt.Println("users:", users)
	return users, nil
}

func AddUser(user models.User) (models.User, error) {
	fmt.Println(user)
	// user.ID = "user1"

	return user, nil
}

func GetUser(id string) (models.User, error) {
	fmt.Println("id : ", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("error parsing ID; ", err)
		return models.User{}, err
	}
	ConnectDB()

	coll := Client.Database("create-app").Collection("users")
	var user models.User
	var res bson.M
	err = coll.FindOne(Ctx, bson.M{"_id": primitive.ObjectID(objID)}).Decode(&res)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found")
	}

	if err != nil {
		fmt.Println("error getting document: ", err)
	}

	bsonBytes, err := bson.Marshal(res)
	if err != nil {
		fmt.Println("error marshalling result: ", err)
		return models.User{}, err
	}

	if err = bson.Unmarshal(bsonBytes, &user); err != nil {
		fmt.Println("error parsing:", err)
		return models.User{}, err
	}

	fmt.Println("user:", user)
	return user, nil

}

func UpdateUser(user models.User) (models.User, error) {
	fmt.Println(user)
	// update the user on db
	// return updated data

	return user, nil
}
