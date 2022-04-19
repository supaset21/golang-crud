package utils

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"gorilla/config"
	"gorilla/models"
)

func CreateBook(title string, stock int) {
	client, ctx, _ := config.MongoDbConnection()
	collection := client.Database(config.DATABASE_NAME).Collection("books")
	book := &models.BooksType{Title: title, Stock: stock}

	if _, err := collection.InsertOne(ctx, book); err != nil {
		// handle error
		//fmt.Println(err)
	}
}

func CheckBookExists(title string) bool {
	result := true
	client, ctx, err := config.MongoDbConnection()
	collection := client.Database(config.DATABASE_NAME).Collection("books")

	var book models.BooksType
	if err = collection.FindOne(ctx, bson.M{"title": title}).Decode(&book); err != nil {
		//fmt.Println(err)
		result = false
	}
	return result
}

func FindBooks() []models.BooksType {
	client, ctx, _ := config.MongoDbConnection()
	collection := client.Database(config.DATABASE_NAME).Collection("books")

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		panic(err)
	}
	defer cur.Close(ctx)

	var books []models.BooksType
	if err = cur.All(ctx, &books); err != nil {
		panic(err)
	}
	return books
}

func FindByTitleBook(title string) (models.BooksType, bool) {
	client, ctx, err := config.MongoDbConnection()
	collection := client.Database(config.DATABASE_NAME).Collection("books")

	var book models.BooksType
	if err = collection.FindOne(ctx, bson.M{"title": title}).Decode(&book); err != nil {
		return models.BooksType{}, false
	}
	return book, true
}

func UpdateStock(title string, stock int) bool {
	isExistsBook := CheckBookExists(title)

	if isExistsBook {
		client, ctx, _ := config.MongoDbConnection()
		collection := client.Database(config.DATABASE_NAME).Collection("books")

		filter := bson.D{{"title", title}}
		update := bson.D{{"$set",
			bson.D{
				{"stock", stock},
			},
		}}

		_, err := collection.UpdateOne(ctx, filter, update)

		if err != nil {
			//fmt.Println(err)
			isExistsBook = false
		}

	}

	return isExistsBook
}

func DeleteBook(title string) bool {
	result := false
	isExistsBook := CheckBookExists(title)

	if isExistsBook {
		client, ctx, _ := config.MongoDbConnection()
		collection := client.Database(config.DATABASE_NAME).Collection("books")
		filter := bson.D{{"title", title}}

		deleteResult, _ := collection.DeleteOne(ctx, filter)

		fmt.Println(deleteResult.DeletedCount)

		if deleteResult.DeletedCount != 0 {
			result = true
		}

	}

	return result
}
