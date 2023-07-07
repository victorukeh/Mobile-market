package models

import (
	"github.com/victorukeh/mobile-market/pkg/v1/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Mint   Type = "mint"
	Normal Type = "normal"
	Both   Type = "both"
)

type CashGroup struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID primitive.ObjectID `bson:"user_id" validate:"required"`
	Type   Type               `json:"type" validate:"oneof=mint normal both,required"`
	Number int64              `json:"number" validate:"required"`
}

var CashGroupCollection *mongo.Collection = database.OpenCollection(database.Client, "cash-group")
