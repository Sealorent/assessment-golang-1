package utils

import (
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(userId string, emailUser string, userName string) (string, error) {

	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(20 * time.Minute).Unix(),
		"context": map[string]interface{}{
			"email":    emailUser,
			"username": userName,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetJWTClaims(tokenString string) (map[string]any, error) {

	secretKey := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid method")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func GetSubFromClaims(claims any) (any, error) {

	mapClaims, ok := claims.(map[string]any)
	if !ok {
		return nil, errors.New("not map")
	}

	sub, ok := mapClaims["sub"]
	if !ok {
		return nil, errors.New("not found")
	}

	return sub, nil
}

func CheckTokenJWTAndReturnSub(ctx *gin.Context) (string, error) {
	claims, exist := ctx.Get("claims")
	if !exist {
		return "", errors.New("unauthorized")
	}

	sub, err := GetSubFromClaims(claims)
	if err != nil {
		return "", errors.New("unauthorized")
	}

	return sub.(string), nil

}
