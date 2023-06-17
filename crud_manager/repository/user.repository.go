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

const collection = "user"

func GetOne(id interface{}) (models.User, error) {
	var user models.User
	objID, _ := primitive.ObjectIDFromHex(fmt.Sprint(id))
	collection, err := database.NewMongoDBConnection(context.Background(), collection)

	if err != nil {
		return user, err
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetByEmail(email string) (models.User, error) {
	var user models.User

	collection, err := database.NewMongoDBConnection(context.Background(), collection)

	if err != nil {
		return user, err
	}
	filter := bson.D{{Key: "email", Value: email}}

	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User

	collection, err := database.NewMongoDBConnection(context.Background(), collection)

	if err != nil {
		return users, err
	}

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return users, err
	}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem models.User
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

func Create(user *models.User) (*models.User, error) {

	var newUser models.User

	collection, err := database.NewMongoDBConnection(context.Background(), collection)

	if err != nil {
		return user, err
	}

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return user, err
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}

	err = collection.FindOne(context.Background(), filter).Decode(&newUser)
	if err != nil {
		return user, err
	}

	return &newUser, nil
}

func Update(user *models.User) (*models.User, error) {
	var updatedUser models.User

	collection, err := database.NewMongoDBConnection(context.Background(), collection)

	if err != nil {
		return user, err
	}
	collection.UpdateByID(context.TODO(), user.Id, user)

	filter := bson.D{{Key: "_id", Value: user.Id}}

	err = collection.FindOne(context.Background(), filter).Decode(&updatedUser)
	if err != nil {
		return user, err
	}

	return &updatedUser, nil
}

func Delete(user *models.User) (*models.User, error) {
	collection, err := database.NewMongoDBConnection(context.Background(), collection)

	if err != nil {
		return user, err
	}
	collection.DeleteOne(context.TODO(), user.Id)

	return user, nil
}
