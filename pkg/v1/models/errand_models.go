package models

import (
	"context"
	"time"

	"github.com/victorukeh/mobile-market/pkg/v1/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Frequency string

const (
	Once         Frequency = "once"
	Weekly       Frequency = "weekly"
	BiWeekly     Frequency = "bi-weekly"
	Monthly      Frequency = "monthly"
	Occasionally Frequency = "occasionally"
)

type Errand struct {
	ID            primitive.ObjectID `bson:"_id"`
	Description   *string            `json:"description" validate:"required"`
	Initiator     primitive.ObjectID `bson:"initiator" validate:"required"`
	ErrandAddress *string            `json:"errand_address" validate:"required"`
	ReturnAddress *string            `json:"return_address" validate:"required"`
	Latitude      float64            `bson:"latitude"`
	Longitude     float64            `bson:"longitude"`
	Frequency     Frequency          `json:"frequency" validate:"oneof=once weekly bi-weekly monthly occasionally,required"`
	Active        bool               `json:"active" validate:"required"`
	Times         []time.Time        `bson:"times"`
	Amount        float64            `json:"amount" validate:"required"`
	People        int64              `json:"people" validate:"required"`
}

var ErrandCollection *mongo.Collection = database.OpenCollection(database.Client, "errand")

func (u *Errand) FindById(id primitive.ObjectID) (Errand, error) {
	value, err := GetErrandById(id)
	if err != nil {
		return value, err
	}
	return value, err
}

func GetErrandById(id primitive.ObjectID) (Errand, error) {
	var errand Errand
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := ErrandCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&errand)
	defer cancel()
	return errand, err
}

func (u *Errand) FindAllErrands(page int64, limit int64) ([]*Errand, error) {
	var results []*Errand
	options := options.Find().SetSort(bson.D{}).SetLimit(limit).SetSkip(page * limit)
	cursor, err := ErrandCollection.Find(context.Background(), bson.D{}, options)
	if err != nil {
		return results, err
	}
	// Iterate over the cursor and decode the documents
	for cursor.Next(context.Background()) {
		var result Errand
		err := cursor.Decode(&result)
		if err != nil {
			return results, err
		}
		results = append(results, &result)
	}
	// Check if there are any errors during iteration
	if err := cursor.Err(); err != nil {
		return results, err
	}
	return results, nil
}

func (U *Errand) CreateErrand(errand Errand) (*mongo.InsertOneResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	errands, err := ErrandCollection.InsertOne(ctx, errand)
	defer cancel()
	return errands, err
}
