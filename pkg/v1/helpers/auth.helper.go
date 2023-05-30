package helper

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/models"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

// SignedDetails
type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Id         primitive.ObjectID
	Role       models.UserRole
	jwt.StandardClaims
}

var key = []byte(SECRET_KEY)

// GenerateAllTokens generates both teh detailed token and refresh token
func GenerateAllTokens(email string, firstName string, lastName string, role models.UserRole, id primitive.ObjectID) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Id:         id,
		Role:       role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(key)

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

// ValidateToken validates the jwt token
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		fmt.Println("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		fmt.Println("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}

func SetSignedCookieOrToken(user models.User) (fiber.Cookie, string) {
	fmt.Println("User: ", user)
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"phone": user.Phone,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}
	fmt.Println("Claims: ", claims)
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	// Send JWT token as cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	return cookie, token
}

func Decode(token string) SignedDetails {
	jwtParts := strings.Split(token, ".")
	payload := jwtParts[1]
	// Decode base64 payload into byte array
	payloadBytes, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		panic(err)
	}

	// Unmarshal JSON payload into struct
	var claims SignedDetails
	err = json.Unmarshal(payloadBytes, &claims)
	if err != nil {
		panic(err)
	}
	return claims
}
