package models

import (
	"context"
	"fmt"
	"time"

	"github.com/victorukeh/mobile-market/pkg/v1/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Currency string
type Type string

const (
	NGN UserRole = "naira"
	USD UserRole = "dollar"
	EUR UserRole = "euro"
	GHS Currency = "cedis"
)

type CashType struct {
	ID       primitive.ObjectID `bson:"_id"`
	Currency *string            `json:"currency" validate:"oneof=naira dollar euro cedis,required"`
	Bill     *int64             `json:"bill" validate:"required" unique:"true"`
	UserID   primitive.ObjectID `bson:"user" validate:"required"`
}

type CashForm struct {
	Status   Status      `json:"status" validate:"oneof=sourcing distributing none,required"`
	CashType []*CashType `json:"cash" validate:"required"`
}

var CashTypeCollection *mongo.Collection = database.OpenCollection(database.Client, "cash-type")

func (u *CashType) GetCashTypes(currency string) ([]*CashType, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var cashTypes []*CashType
	options := options.Find()
	cursor, err := CashTypeCollection.Find(ctx, bson.M{"currency": currency}, options)
	if err != nil {
		defer cancel()
		return cashTypes, err
	}

	for cursor.Next(ctx) {
		var cashType CashType
		err := cursor.Decode(&cashType)
		if err != nil {
			defer cancel()
			return cashTypes, err
		}
		cashTypes = append(cashTypes, &cashType)
	}

	if err := cursor.Err(); err != nil {
		defer cancel()
		return cashTypes, err
	}
	defer cancel()
	return cashTypes, err

}

func (u *CashType) CreateCashType(cashType CashType) (CashType, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	value, InsertErr := CashTypeCollection.InsertOne(ctx, cashType)
	err := CashTypeCollection.FindOne(ctx, bson.M{"_id": value}).Decode(&cashType)
	if err != nil {
		err = InsertErr
	}
	defer cancel()
	return cashType, err
}

func (u *CashType) FindOneCashType(cashType CashType) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.D{{Key: "bill", Value: cashType.Bill}, {Key: "currency", Value: cashType.Currency}}
	fmt.Println(*cashType.Bill, *cashType.Currency)
	err := CashTypeCollection.FindOne(ctx, filter).Decode(&cashType)
	defer cancel()
	return err
}

func (u *CashType) FindCashTypeById(id primitive.ObjectID) error {
	fmt.Println("ID: ", id)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var cashType CashType
	err := CashTypeCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&cashType)

	defer cancel()
	return err
}
