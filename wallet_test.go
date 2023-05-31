package main

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/victorukeh/mobile-market/pkg/v1/models"
)

func TestFindById(t *testing.T) {
	var wallet models.Wallet
	useWallet := models.Wallet{ID: primitive.NewObjectID(), Balance: 30}
	// Insert test data into collection
	_, err := models.WalletCollection.InsertOne(context.Background(), useWallet)
	if err != nil {
		t.Fatal(err)
	}
	// // Clean up test data
	defer func() {
		_, err := models.WalletCollection.DeleteOne(context.Background(), bson.M{"_id": useWallet.ID})
		if err != nil {
			t.Fatal(err)
		}

	}()
	// Positive test case
	result, err := wallet.FindById(useWallet.ID)
	if err != nil {
		t.Errorf("Expected nil error, but got %v", err)
	}

	// Negative test case
	invalidId := primitive.NewObjectID()
	result, err = wallet.FindById(invalidId)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	if result.ID == invalidId {
		t.Errorf("Expected result to have ID %v, but got %v", invalidId, result.ID)
	}
}

func TestFindAll(t *testing.T) {
	// Set up test data
	var walletModel models.Wallet
	wallets := []*models.Wallet{
		{ID: primitive.NewObjectID(), Balance: 10},
		{ID: primitive.NewObjectID(), Balance: 20},
		{ID: primitive.NewObjectID(), Balance: 30},
	}
	// Insert test data into collection
	for _, wallet := range wallets {
		_, err := models.WalletCollection.InsertOne(context.Background(), wallet)
		if err != nil {
			t.Fatal(err)
		}
	}
	// // Clean up test data
	defer func() {
		for _, wallet := range wallets {
			_, err := models.WalletCollection.DeleteOne(context.Background(), bson.M{"_id": wallet.ID})
			if err != nil {
				t.Fatal(err)
			}
		}

	}()

	// Positive test cases
	t.Run("Valid page and limit parameters", func(t *testing.T) {
		page := int64(0)
		limit := int64(2)
		results, err := walletModel.FindAll(page, limit)
		if err != nil {
			t.Fatal(err)
		}
		if len(results) != 2 {
			t.Errorf("Expected %d results, but got %d", 2, len(results))
		}
	})
	t.Run("Page and limit parameters set to 0", func(t *testing.T) {
		page := int64(0)
		limit := int64(0)
		results, err := walletModel.FindAll(page, limit)
		if err != nil {
			t.Fatal(err)
		}
		if len(results) != 0 {
			t.Errorf("Expected %d results, but got %d", 0, len(results))
		}
	})
	t.Run("Page parameter greater than total number of documents", func(t *testing.T) {
		page := int64(5)
		limit := int64(2)
		results, err := walletModel.FindAll(page, limit)
		if err != nil {
			t.Fatal(err)
		}
		if len(results) != 0 {
			t.Errorf("Expected %d results, but got %d", 0, len(results))
		}
	})
	// Negative test cases
	t.Run("Invalid page parameter", func(t *testing.T) {
		page := int64(-1)
		limit := int64(2)
		_, err := walletModel.FindAll(page, limit)
		if err == nil {
			t.Error("Expected an error, but got nil")
		}
	})
	t.Run("Invalid limit parameter", func(t *testing.T) {
		page := int64(0)
		limit := int64(-1)
		_, err := walletModel.FindAll(page, limit)
		if err == nil {
			t.Error("Expected an error, but got nil")
		}
	})
	// t.Run("Invalid collection name", func(t *testing.T) {
	// 	page := int64(0)
	// 	limit := int64(2)
	// 	_, err := FindAll(page, limit, "invalid_collection_name")
	// 	if err == nil {
	// 		t.Error("Expected an error, but got nil")
	// 	}
	// })
}

func ConvertToObjectID(id string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return objID, nil
}
