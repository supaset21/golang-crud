package utils

import (
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

	data, err := collection.Find(ctx, bson.D{})

	if err != nil {
		panic(err)
	}
	defer data.Close(ctx)

	var books []models.BooksType
	if err = data.All(ctx, &books); err != nil {
		panic(err)
	}
	return books
}
