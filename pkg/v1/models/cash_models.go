package models

import (
	"github.com/victorukeh/mobile-market/pkg/v1/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Currency string
type Type string

const (
	NGN UserRole = "naira"
	USD UserRole = "dollar"
	EUR UserRole = "euro"
	GHS Currency = "cedis"
)

const (
	Mint   Type = "mint"
	Normal Type = "normal"
	Both   Type = "both"
)

type Cash struct {
	ID       primitive.ObjectID `bson:"id"`
	Currency *string            `json:"currency" validate:"oneof=naira dollar euro cedis,required"`
	Bill     *int64             `json:"bill" validate:"required"`
	Amount   *string            `json:"amount" validate:"required"`
	UserID   primitive.ObjectID `bson:"user_id" validate:"required"`
	Type     Type               `json:"type" validate:"oneof=mint normal both,required"`
}

type CashForm struct {
	Status Status  `json:"status" validate:"oneof=sourcing distributing none,required"`
	Cash   []*Cash `json:"cash" validate:"required"`
}

var CashCollection *mongo.Collection = database.OpenCollection(database.Client, "cash")
