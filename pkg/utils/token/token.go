package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("tUrb0-Bp0**")

type Claims struct {
	Email  string    `json:"email"`
	UserId string    `json:"userId"`
	Name   string    `json:"name"`
	Role   []*string `json:"role"`
	jwt.StandardClaims
}

func New(email string, userID string, name string, role []*string) (*string, error) {

	expirationTime := time.Now().Add((24 * 30) * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email:  email,
		UserId: userID,
		Name:   name,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return nil, err
	}
	return &tokenString, nil

}

func Decode(bearerToken string) (*Claims, error) {
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//return &Claims{Email: fmt.Sprintf("%v", claims["email"]), Role: token.Claims.Role , Name: fmt.Sprintf("%v", claims["name"]), UserId: fmt.Sprintf("%v", claims["userId"])}, err
		var roles []*string
		for _, v := range claims["role"].([]interface{}) {
			valStr := v.(string)
			roles = append(roles, &valStr)
		}
		return &Claims{Email: fmt.Sprintf("%v", claims["email"]), Role: roles, Name: fmt.Sprintf("%v", claims["name"]), UserId: fmt.Sprintf("%v", claims["userId"])}, err
	} else {
		return nil, err
	}
}