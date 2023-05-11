package db

import (
	"fmt"
	"go-practice/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPosts() ([]models.Post, error) {
	posts := []models.Post{}
	ConnectDB()

	coll := Client.Database("create-app").Collection("posts")
	cursor, err := coll.Find(Ctx, bson.D{})
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found")
		return []models.Post{}, err
	}

	if err != nil {
		fmt.Println("error getting document: ", err)
		return []models.Post{}, err
	}

	if err = cursor.All(Ctx, &posts); err != nil {
		fmt.Println("error parsing: ", err)
		return []models.Post{}, err

	}
	return posts, nil

}

func AddPost(post models.Post) (models.Post, error) {
	// add to db
	// return the added post including the post ID.

	post.ID = primitive.NewObjectID()

	ConnectDB()
	coll := Client.Database("create-app").Collection("posts")
	insertResult, err := coll.InsertOne(Ctx, post)
	if err != nil {
		fmt.Println("error adding post to db")
		return models.Post{}, err
	}

	postId := insertResult.InsertedID
	postIdObj := postId.(primitive.ObjectID)
	post.ID = postIdObj

	return post, nil
}

func GetPost(id string) (models.Post, error) {

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("error parsing id: ", err)
		return models.Post{}, err
	}

	ConnectDB()

	coll := Client.Database("create-app").Collection("posts")
	post := models.Post{}
	var res bson.M
	err = coll.FindOne(Ctx, bson.M{"_id": primitive.ObjectID(objId)}).Decode(&res)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No document was found")
		return models.Post{}, err
	}

	if err != nil {
		fmt.Println("error getting document: ", err)
		return models.Post{}, err
	}

	bsonByte, err := bson.Marshal(res)
	if err != nil {
		fmt.Println("error marshalling result: ", err)
		return models.Post{}, err
	}

	if err = bson.Unmarshal(bsonByte, &post); err != nil {
		fmt.Println("error parsing:", err)
		return models.Post{}, err
	}

	fmt.Println("post: ", post)
	return post, nil
}

func UpdatePost(post models.Post) (models.Post, error) {
	// update the post on db
	// return updated data

	postByte, err := bson.Marshal(post)
	if err != nil {
		fmt.Println("error marshalling post data: ", err)
		return models.Post{}, err
	}

	var update bson.M
	if err = bson.Unmarshal(postByte, &update); err != nil {
		fmt.Println("error unmarshalling: ", err)
		return models.Post{}, err
	}

	ConnectDB()
	coll := Client.Database("create-app").Collection("posts")
	_, err = coll.UpdateOne(Ctx, bson.M{"_id": post.ID}, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		fmt.Println("error updating post: ", err)
		return models.Post{}, err
	}
	return post, nil
}
