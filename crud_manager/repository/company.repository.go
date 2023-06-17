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

const collectionCompany = "company"

func GetOneCompany(id interface{}) (models.Company, error) {
	var company models.Company
	objID, _ := primitive.ObjectIDFromHex(fmt.Sprint(id))
	collection, err := database.NewMongoDBConnection(context.Background(), collectionCompany)

	if err != nil {
		return company, err
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&company)
	if err != nil {
		return company, err
	}

	return company, nil
}

func GetByEmailCompany(email string) (models.Company, error) {
	var company models.Company

	collection, err := database.NewMongoDBConnection(context.Background(), collectionCompany)

	if err != nil {
		return company, err
	}
	filter := bson.D{{Key: "email", Value: email}}

	err = collection.FindOne(context.Background(), filter).Decode(&company)
	if err != nil {
		return company, err
	}

	return company, nil
}

func GetAllCompanies() ([]models.Company, error) {
	var companys []models.Company

	collection, err := database.NewMongoDBConnection(context.Background(), collectionCompany)

	if err != nil {
		return companys, err
	}

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return companys, err
	}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem models.Company
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		companys = append(companys, elem)

	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return companys, nil
}

func CreateCompany(company *models.Company) (*models.Company, error) {

	var newCompany models.Company

	collection, err := database.NewMongoDBConnection(context.Background(), collectionCompany)

	if err != nil {
		return company, err
	}

	result, err := collection.InsertOne(context.Background(), company)
	if err != nil {
		return company, err
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}

	err = collection.FindOne(context.Background(), filter).Decode(&newCompany)
	if err != nil {
		return company, err
	}

	return &newCompany, nil
}

func UpdateCompany(company *models.Company) (*models.Company, error) {
	var updatedcompany models.Company

	collection, err := database.NewMongoDBConnection(context.Background(), collectionCompany)

	if err != nil {
		return company, err
	}
	collection.UpdateByID(context.TODO(), company.Id, company)

	filter := bson.D{{Key: "_id", Value: company.Id}}

	err = collection.FindOne(context.Background(), filter).Decode(&updatedcompany)
	if err != nil {
		return company, err
	}

	return &updatedcompany, nil
}

func DeleteCompany(company *models.Company) (*models.Company, error) {
	collection, err := database.NewMongoDBConnection(context.Background(), collectionCompany)

	if err != nil {
		return company, err
	}
	collection.DeleteOne(context.TODO(), company.Id)

	return company, nil
}
