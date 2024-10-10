package auth

import (
	"encoding/base64"
	"fmt"

	"github.com/golang-jwt/jwt"
)

func ValidateJWT(tokenString string, publicKey string) (bool, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return false, fmt.Errorf("error decoding public key: %v", err)
	}
	// Parse the public key
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(decodedPublicKey))
	if err != nil {
		return false, fmt.Errorf("error parsing public key: %v", err)
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return false, fmt.Errorf("error parsing token: %v", err)
	}

	// Validate the token
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}

	return false, nil
}
