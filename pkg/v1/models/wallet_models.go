package models

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/victorukeh/mobile-market/pkg/v1/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Wallet struct {
	ID      primitive.ObjectID `bson:"_id"`
	UserID  primitive.ObjectID `bson:"user_id" validate:"required" unique:"true"`
	Balance float64            `json:"balance"`
}

var WalletCollection *mongo.Collection = database.OpenCollection(database.Client, "wallet")

func (u *Wallet) FindById(id primitive.ObjectID) (Wallet, error) {
	value, err := GetWalletById(id)
	if err != nil {
		return value, err
	}
	return value, err
}

func (u *Wallet) FindAll(page int64, limit int64) ([]*Wallet, error) {
	var results []*Wallet

	if limit < 0 || page < 0 {
		return results, errors.New("invalid page or limit")
	}
	if limit == 0 && page == 0 {
		return results, nil
	}
	options := options.Find().SetSort(bson.D{}).SetLimit(limit).SetSkip(page * limit)
	cursor, err := WalletCollection.Find(context.Background(), bson.D{}, options)
	if err != nil {
		return results, err
	}
	// Iterate over the cursor and decode the documents
	for cursor.Next(context.Background()) {
		var result Wallet
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
	return results, err
}

func (u *Wallet) FindByIdAndDelete(id primitive.ObjectID) (*mongo.DeleteResult, Wallet, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	wallet, _ := GetWalletById(id)
	filter := bson.M{"_id": id}
	result, err := WalletCollection.DeleteOne(ctx, filter)
	defer cancel()
	return result, wallet, err
}

func (u *Wallet) FindByEmail(email string, wallet Wallet) (Wallet, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	err := WalletCollection.FindOne(ctx, bson.M{"email": email}).Decode(&wallet)
	if err != nil {
		fmt.Println(err)
	}
	defer cancel()
	return wallet, err
}

func (u *Wallet) Create(wallet Wallet) (Wallet, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := WalletCollection.FindOne(ctx, bson.M{"user_id": wallet.UserID}).Decode(&wallet)
	fmt.Println("Error", err)
	if err == nil {
		defer cancel()
		return wallet, err
	}
	_, err = WalletCollection.InsertOne(ctx, wallet)
	fmt.Println("Error2", err)
	defer cancel()
	return wallet, err
}

func (u *Wallet) CreateMany(wallets []Wallet) (*mongo.InsertManyResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	items, err := WalletCollection.InsertMany(ctx, sliceWalletsToInterface(wallets))
	defer cancel()
	return items, err
}

func (u *Wallet) CountWallets(field string, value string) (int64, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	count, err := WalletCollection.CountDocuments(ctx, bson.M{field: value})
	defer cancel()
	return count, err
}

func (u *Wallet) FindByUserID(id primitive.ObjectID) (Wallet, error) {
	var wallet Wallet
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := WalletCollection.FindOne(ctx, bson.M{"user_id": id}).Decode(&wallet)
	fmt.Println("Errpr", err, id)
	defer cancel()
	return wallet, err
}

func sliceWalletsToInterface(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("sliceToInterface() called with non-slice value")
	}
	result := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		result[i] = s.Index(i).Interface()
	}
	return result
}

func GetWalletById(id primitive.ObjectID) (Wallet, error) {
	var wallet Wallet
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := WalletCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&wallet)
	defer cancel()
	return wallet, err
}

func (u *Wallet) Count(filter options.ArrayFilters) (int64, error) {
	ctx := context.Background()

	count, err := WalletCollection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}
