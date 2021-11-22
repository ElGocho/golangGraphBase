package helpers

import (
	"fmt"
	"time"
	"context"
	"github.com/dgrijalva/jwt-go"

	"sa_web_service/internal/models/consts"
)

var (
	TOKEN_TIME_LIFE = 24
)

func GenerateToken(data jwt.MapClaims) (string, error){
	if data["exp"] == nil {
		data["exp"] = time.Now().Add(time.Hour * 24).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,data)

	secretKey := []byte(consts.TokenLoginKey)

	return token.SignedString(secretKey)
}

func ParseToken(signedString string) (jwt.MapClaims, error) {
	if signedString == ""{
		return nil, fmt.Errorf("SignedString empty")
	}

	token, err := jwt.Parse(signedString, func (token *jwt.Token) (interface{}, error){
		if _,OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(consts.TokenLoginKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	if claims,OK := token.Claims.(jwt.MapClaims); OK{
		return claims, nil
	}

	return nil, fmt.Errorf("Error to parse token")
}

func GetLoginClaims(ctx context.Context) (jwt.MapClaims, error){
	token, OK := ctx.Value(consts.AuthorizationKey).(string); 

	if !OK{
		return nil, fmt.Errorf("Error to found claims")
	}

	claims, err := ParseToken(token)

	return claims, err
}

func ValidAuthToken(ctx context.Context) error{
	token, OK := ctx.Value(consts.AuthorizationKey).(string); 

	if !OK{
		return fmt.Errorf("Error to found claims")
	}

	_, err := ParseToken(token)

	return err
}

