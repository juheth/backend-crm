package common

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Expira en 72 horas

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Asegurarse de que el token es válido y retornar los claims correctos
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token inválido")
}

func RefreshToken(tokenString string) (string, int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", 0, fmt.Errorf("token inválido: %v", err)
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return "", 0, fmt.Errorf("token inválido o claims inválidos")
	}

	userID := claims.ID

	expiration := claims.ExpiresAt
	if time.Until(time.Unix(expiration, 0)) < time.Hour*24 {
		newToken, err := GenerateToken(userID)
		return newToken, userID, err
	}
	return tokenString, userID, nil
}
