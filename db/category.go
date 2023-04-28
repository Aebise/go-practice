package db

import (
	"fmt"
	"go-practice/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCategories() (categories []models.Category, err error) {
	categories = []models.Category{}

	ConnectDB()
	coll := Client.Database("create-app").Collection("categories")
	cursor, err := coll.Find(Ctx, bson.D{})
	if err == mongo.ErrNoDocuments {
		fmt.Println("No document was found")
		return []models.Category{}, err
	}

	if err != nil {
		fmt.Println("error getting document: ", err)
		return []models.Category{}, err
	}

	if err = cursor.All(Ctx, &categories); err != nil {
		fmt.Println("error parsing: ", err)
		return []models.Category{}, err
	}

	return categories, nil

}

func AddCategory(category models.Category) (models.Category, error) {
	// add to db
	// return the added category including the category ID.

	category.ID = primitive.NewObjectID()
	ConnectDB()
	coll := Client.Database("create-app").Collection("categories")
	insertResult, err := coll.InsertOne(Ctx, &category)
	if err != nil {
		fmt.Println("error adding category to db: ", err)
		return models.Category{}, err
	}

	categoryId := insertResult.InsertedID
	categoryIdObj := categoryId.(primitive.ObjectID)
	category.ID = categoryIdObj

	return category, nil
}

func GetCategory(id string) (models.Category, error) {

	category := models.Category{}
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("error parsing id: ", err)
		return models.Category{}, err
	}

	ConnectDB()
	coll := Client.Database("create-app").Collection("categories")
	var res bson.M
	err = coll.FindOne(Ctx, bson.M{"_id": primitive.ObjectID(objId)}).Decode(&res)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No document was found: ", err)
		return models.Category{}, err
	}

	if err != nil {
		fmt.Println("error getting document: ", err)
		return models.Category{}, err
	}

	bsonByte, err := bson.Marshal(res)
	if err != nil {
		fmt.Println("error marshalling result: ", err)
		return models.Category{}, err
	}

	if err = bson.Unmarshal(bsonByte, &category); err != nil {
		fmt.Println("error parsing: ", err)
		return models.Category{}, err
	}

	fmt.Println("category:", category)

	return category, nil
}

func UpdateCategory(category models.Category) (models.Category, error) {
	// update the category on db
	// return updated data

	categoryByte, err := bson.Marshal(category)
	if err != nil {
		fmt.Println("error marshalling category data: ", err)
		return models.Category{}, err
	}

	var update bson.M
	if err = bson.Unmarshal(categoryByte, &update); err != nil {
		fmt.Println("error unmarshalling: ", err)
		return models.Category{}, err
	}

	ConnectDB()
	coll := Client.Database("create-app").Collection("categories")
	_, err = coll.UpdateOne(Ctx, bson.M{"_id": category.ID}, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		fmt.Println("error updating category: ", err)
		return models.Category{}, err
	}

	return category, nil
}
