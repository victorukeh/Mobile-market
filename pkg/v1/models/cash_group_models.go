package models

import (
	"context"
	"time"

	"github.com/victorukeh/mobile-market/pkg/v1/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Mint   Type = "mint"
	Normal Type = "normal"
	Both   Type = "both"
)

type CashData struct {
	CashGroup []CashGroup `json:"cashGroup"`
}

type CashGroup struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserID     primitive.ObjectID `bson:"user" validate:"required"`
	CashTypeID primitive.ObjectID `bson:"cashType" validate:"required"`
	Type       Type               `json:"type" validate:"oneof=mint normal both,required"`
	Number     int64              `json:"number" validate:"required"`
}

var CashGroupCollection *mongo.Collection = database.OpenCollection(database.Client, "cash-group")

func (u *CashGroup) CreateCashGroup(cashGroup CashGroup) (CashGroup, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	value, InsertErr := CashGroupCollection.InsertOne(ctx, cashGroup)
	err := CashGroupCollection.FindOne(ctx, bson.M{"_id": value}).Decode(&cashGroup)
	if err != nil {
		err = InsertErr
	}
	defer cancel()
	return cashGroup, err
}

func (u *CashGroup) FindCashGroup(cashGroup CashGroup) (*mongo.SingleResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.D{{Key: "type", Value: cashGroup.Type}, {Key: "cashType", Value: cashGroup.CashTypeID}}
	value := CashGroupCollection.FindOne(ctx, filter)
	err := CashGroupCollection.FindOne(ctx, filter).Decode(&cashGroup)
	defer cancel()
	return value, err
}

func (u *CashGroup) UpdateCashGroup(id primitive.ObjectID, number int64) error {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "number", Value: number}}}}
	_, err := CashGroupCollection.UpdateOne(context.TODO(), filter, update)
	return err
}
