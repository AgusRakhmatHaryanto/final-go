package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(ttl time.Duration, payload interface{}, role string, id int, secretJWTKey string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    now := time.Now().UTC()
    claims := token.Claims.(jwt.MapClaims)

    claims["id"] = id
    claims["role"] = role
    claims["sub"] = payload

    claims["exp"] = now.Add(ttl).Unix()
    claims["iat"] = now.Unix()
    claims["nbf"] = now.Unix()

    tokenString, err := token.SignedString([]byte(secretJWTKey))

    if err != nil {
        return "", fmt.Errorf("generating JWT Token failed: %w", err)
    }

    return tokenString, nil
}



func ValidateToken(token string, signedJWTKey string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims["sub"], nil
}

func ExtractToken(c *gin.Context) (string, string, error) {

	user, exist := c.Get("currentUser")
	if !exist {
		return "", "", errors.New("invalid token")
	}

	claims := user.(gin.H)

	id, ok := claims["id"].(string)
	if !ok {
		return "", "", errors.New("invalid token")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", "", errors.New("invalid token")
	}

	return id, role, nil
}
