package tokenmanager

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var InvalidToken = errors.New("invalid token")

type TokenController struct {
	tokenTTL  time.Duration
	secretKey []byte
}

type Claims struct {
	jwt.RegisteredClaims
	UserUUID string
}

func New(exp time.Duration, key []byte) *TokenController {
	return &TokenController{
		tokenTTL:  exp,
		secretKey: key,
	}
}

func (tk *TokenController) CreateToken(uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tk.tokenTTL)),
		},
		UserUUID: uuid,
	})

	tokenString, err := token.SignedString(tk.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (tk *TokenController) GetUserUUID(tokenString string) (string, error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return tk.secretKey, nil
	})

	if err != nil {
		return "", InvalidToken
	}

	if !token.Valid {

		return "", InvalidToken
	}

	return claims.UserUUID, nil
}
