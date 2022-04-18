package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type BooksType struct {
	Title     string    `json:"title" bson:"title"`
	Stock     int       `json:"stock" bson:"stock"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func (book *BooksType) MarshalBSON() ([]byte, error) {
	now := time.Now()
	if book.CreatedAt.IsZero() {
		book.CreatedAt = now
	}
	book.UpdatedAt = now

	type my BooksType
	return bson.Marshal((*my)(book))
}
