package models

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"

	"github.com/victorukeh/mobile-market/pkg/v1/database"
)

type UserRole string

const (
	Admin    UserRole = "admin"
	Consumer UserRole = "user"
	Store    UserRole = "store"
)

type User struct {
	ID                 primitive.ObjectID `bson:"_id"`
	First_name         *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name          *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password           *string            `json:"password" validate:"required,min=6"`
	Email              *string            `json:"email" validate:"email,required"`
	Country            *string            `json:"country" validate:"required"`
	Phone              *string            `json:"phone" validate:"required"`
	Role               UserRole           `json:"role" validate:"oneof= admin user store","required"`
	Confirmation_token *string            `json:"confirmation_token"`
	Created_at         time.Time          `json:"created_at"`
	Updated_at         time.Time          `json:"updated_at"`
	// User_id       string             `json:"user_id"`
}

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func (u *User) HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

// VerifyPassword checks the input password while verifying it with the password in the DB.
func (u *User) VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true

	if err != nil {
		check = false
	}

	return check
}

func (u *User) FindById(id primitive.ObjectID) (User, error) {
	value, err := GetById(id)
	if err != nil {
		return value, err
	}
	return value, err
}

func (u *User) FindAll(page int64, limit int64) ([]*User, error) {
	var results []*User
	options := options.Find().SetSort(bson.D{}).SetLimit(limit).SetSkip(page * limit)
	cursor, err := UserCollection.Find(context.Background(), bson.D{}, options)
	if err != nil {
		return results, err
	}
	// Iterate over the cursor and decode the documents
	for cursor.Next(context.Background()) {
		var result User
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

func (u *User) FindByIdAndDelete(id primitive.ObjectID) (*mongo.DeleteResult, User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	user, _ := GetById(id)
	filter := bson.M{"_id": id}
	result, err := UserCollection.DeleteOne(ctx, filter)
	defer cancel()
	return result, user, err
}

func (u *User) FindOne(user User) *mongo.SingleResult {
	result := GetOne(user)
	return result
}

func (u *User) FindByEmail(email string, user User) (User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	err := UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	fmt.Println("Check5")
	if err != nil {
		fmt.Println(err)
	}
	defer cancel()
	return user, err
}

func (u *User) Create(user User) (User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	value, InsertErr := UserCollection.InsertOne(ctx, user)
	err := UserCollection.FindOne(ctx, bson.M{"_id": value}).Decode(&user)
	if err != nil {
		err = InsertErr
	}
	defer cancel()
	return user, err
}

func (u *User) CreateMany(users []User) (*mongo.InsertManyResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	items, err := UserCollection.InsertMany(ctx, sliceToInterface(users))
	defer cancel()
	return items, err
}

func (u *User) CountUsers(field string, value string) (int64, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	count, err := UserCollection.CountDocuments(ctx, bson.M{field: value})
	defer cancel()
	return count, err
}

func sliceToInterface(slice interface{}) []interface{} {
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

func GetById(id primitive.ObjectID) (User, error) {
	var user User
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := UserCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	defer cancel()
	return user, err
}

func GetOne(user User) *mongo.SingleResult {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	result := UserCollection.FindOne(ctx, bson.M{"first_name": user.First_name,
		"last_name": user.Last_name,
		"email":     user.Email,
		"country":   user.Country,
		"phone":     user.Phone})
	defer cancel()
	return result
}
