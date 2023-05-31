package tests

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/victorukeh/mobile-market/pkg/v1/models"
)

func TestFindById(t *testing.T) {
	var wallet models.Wallet
	// Positive test case
	id, _ := ConvertToObjectID("6477147d2510ccd0f25fe14f")
	result, err := wallet.FindById(id)
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

func ConvertToObjectID(id string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return objID, nil
}
