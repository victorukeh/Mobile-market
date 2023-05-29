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
			return key, nil
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

	fmt.Println(claims.Role)
	return claims, msg
}

func SetSignedCookieOrToken(user models.User) (fiber.Cookie, string) {
	token := jwt.New(jwt.SigningMethodHS256)
	// duration, _ := time.ParseDuration(JWT_EXPIRE)
	// 	// claims := token.Claims.(jwt.MapClaims)

	// 	// setClaims := jwt.MapClaims{
	// 	// 	"id":    user.ID,
	// 	// 	"email": user.Email,
	// 	// 	"role":  user.Role,
	// 	// 	"phone": user.Phone,
	// 	// 	"exp":   time.Now().Add(duration).Unix(), // Expires in 24 hours
	// 	// }
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["phone"] = user.Phone
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Expires in 24 hours
	signedToken, _ := token.SignedString(key)

	check, _ := ValidateToken(signedToken)

	fmt.Println("123: ", check, signedToken)
	// claims := jwt.MapClaims{
	// 	"id":    user.ID,
	// 	"email": user.Email,
	// 	"role":  user.Role,
	// 	"phone": user.Phone,
	// 	"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	// }
	// token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	// // Send JWT token as cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	return cookie, signedToken
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

// package helper

// import (
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/victorukeh/mobile-market/pkg/v1/models"

// 	jwt "github.com/dgrijalva/jwt-go"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// var SECRET_KEY string = os.Getenv("SECRET_KEY")
// var JWT_EXPIRE string = os.Getenv("JWT_EXPIRE")

// // SignedDetails
// type SignedDetails struct {
// 	Email      string
// 	First_name string
// 	Last_name  string
// 	Id         primitive.ObjectID
// 	Role       models.UserRole
// 	jwt.StandardClaims
// }

// var key = []byte(SECRET_KEY)

// // GenerateAllTokens generates both teh detailed token and refresh token
// func GenerateAllTokens(email string, firstName string, lastName string, role models.UserRole, id primitive.ObjectID) (signedToken string, signedRefreshToken string, err error) {
// 	claims := &SignedDetails{
// 		Email:      email,
// 		First_name: firstName,
// 		Last_name:  lastName,
// 		Id:         id,
// 		Role:       role,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
// 		},
// 	}

// 	refreshClaims := &SignedDetails{
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
// 		},
// 	}

// 	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
// 	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(key)

// 	if err != nil {
// 		log.Panic(err)
// 		return
// 	}

// 	return token, refreshToken, err
// }

// // ValidateToken validates the jwt token
// func ValidateToken(signedToken string) (claims jwt.Claims, msg string) {
// 	token, err := jwt.Parse(
// 		signedToken,
// 		// &SignedDetails{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return key, nil
// 		},
// 	)

// 	if err != nil {
// 		msg = err.Error()
// 		return
// 	}

// 	if !token.Valid {
// 		fmt.Println("Invalid JWT token")
// 		return
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		fmt.Println("the token is invalid")
// 		msg = err.Error()
// 		return
// 	}
// 	name, ok := claims["name"].(string)
//     if !ok {
//         fmt.Println("Failed to extract name from JWT claims")
//         return
//     }
// 	// if claims.ExpiresAt < time.Now().Local().Unix() {
// 	// 	fmt.Println("token is expired")
// 	// 	msg = err.Error()
// 	// 	return
// 	// }

// 	return claims, msg
// }

// func SetSignedCookieOrToken(user models.User) (fiber.Cookie, string) {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	duration, _ := time.ParseDuration(JWT_EXPIRE)
// 	// claims := token.Claims.(jwt.MapClaims)

// 	// setClaims := jwt.MapClaims{
// 	// 	"id":    user.ID,
// 	// 	"email": user.Email,
// 	// 	"role":  user.Role,
// 	// 	"phone": user.Phone,
// 	// 	"exp":   time.Now().Add(duration).Unix(), // Expires in 24 hours
// 	// }
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["id"] = user.ID
// 	claims["email"] = user.Email
// 	claims["role"] = user.Role
// 	claims["phone"] = user.Phone
// 	claims["exp"] = time.Now().Add(duration).Unix() // Expires in 24 hours
// 	signedToken, _ := token.SignedString(key)
// 	check, _ := ValidateToken(signedToken)
// 	fmt.Println("Check: ", check)

// 	// Send JWT token as cookie
// 	cookie := fiber.Cookie{
// 		Name:     "jwt",
// 		Value:    signedToken,
// 		Expires:  time.Now().Add(time.Hour * 24),
// 		HTTPOnly: true,
// 	}
// 	return cookie, signedToken
// }

// func Decode(token string) SignedDetails {
// 	jwtParts := strings.Split(token, ".")
// 	payload := jwtParts[1]
// 	// Decode base64 payload into byte array
// 	payloadBytes, err := base64.RawURLEncoding.DecodeString(payload)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Unmarshal JSON payload into struct
// 	var claims SignedDetails
// 	err = json.Unmarshal(payloadBytes, &claims)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return claims
// }
