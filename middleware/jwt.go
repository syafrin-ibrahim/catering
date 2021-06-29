package middleware

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	///"github.com/labstack/echo/v4/middleware"
)

// var isAdmin = middleware.JWTConfig(middleware.JWTConfig{
// 	SigningKey: []byte(constant.SECRET_JWT),
// })
const SECRET_JWT = "1234567"

func CreateToken(userId int, isadmin bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = isadmin
	claims["id_user"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET_JWT))
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_JWT), nil
	})

	if err != nil {
		return token, nil
	}

	return token, err

}
