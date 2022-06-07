package helper

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "secret"

func GeneratedToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(SECRET_KEY))
	return signedToken

}

func VerifyToken(tokenString string) (interface{}, error) {
	errResp := errors.New("token invalid")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResp
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResp
	}
	return token.Claims.(jwt.MapClaims), nil

}
