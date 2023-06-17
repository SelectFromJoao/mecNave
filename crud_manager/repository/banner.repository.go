package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mecnave.com/mod/crud_manager/database"
	"mecnave.com/mod/crud_manager/models"
)

const collectionBanner = "banner"

func GetOneBanner(id interface{}) (models.Banner, error) {
	var banner models.Banner
	objID, _ := primitive.ObjectIDFromHex(fmt.Sprint(id))
	collection, err := database.NewMongoDBConnection(context.Background(), collectionBanner)

	if err != nil {
		return banner, err
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&banner)
	if err != nil {
		return banner, err
	}

	return banner, nil
}

func GetAllBanners() ([]models.Banner, error) {
	var users []models.Banner

	collection, err := database.NewMongoDBConnection(context.Background(), collectionBanner)

	if err != nil {
		return users, err
	}

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return users, err
	}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem models.Banner
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, elem)

	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func CreateBanner(banner *models.Banner) (*models.Banner, error) {

	var newUser models.Banner

	collection, err := database.NewMongoDBConnection(context.Background(), collectionBanner)

	if err != nil {
		return banner, err
	}

	result, err := collection.InsertOne(context.Background(), banner)
	if err != nil {
		return banner, err
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}

	err = collection.FindOne(context.Background(), filter).Decode(&newUser)
	if err != nil {
		return banner, err
	}

	return &newUser, nil
}

func DeleteBanner(banner *models.Banner) (*models.Banner, error) {
	collection, err := database.NewMongoDBConnection(context.Background(), collectionBanner)

	if err != nil {
		return banner, err
	}
	collection.DeleteOne(context.TODO(), banner.Id)

	return banner, nil
}
