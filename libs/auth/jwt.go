package auth

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Avatar        string `json:"avatar"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"emailVerified"`
	Name          string `json:"name"`
	NickName      string `json:"nickName"`
}

type Claims struct {
	User UserClaims `json:"user"`
	Exp  int        `json:"exp"`
}

func ValidateJWT(tokenString string, publicKey string) (claims Claims, err error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return claims, fmt.Errorf("error decoding public key: %v", err)
	}
	// Parse the public key
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(decodedPublicKey))
	if err != nil {
		return claims, fmt.Errorf("error parsing public key: %v", err)
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return claims, fmt.Errorf("error parsing token: %v", err)
	}

	// Validate the token
	if tokenClaims, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		claims = convertClaims(tokenClaims)

		return claims, errors.New("invalid token")
	}

	return claims, nil
}

func convertClaims(tokenClaims map[string]interface{}) Claims {
	user, ok := tokenClaims["user"].(map[string]interface{})
	if !ok {
		return Claims{}
	}

	avatar, ok := user["avatar"].(string)
	if !ok {
		avatar = ""
	}

	email, ok := user["email"].(string)
	if !ok {
		email = ""
	}

	emailVerified, ok := user["email_verified"].(bool)
	if !ok {
		emailVerified = false
	}

	name, ok := user["name"].(string)
	if !ok {
		name = ""
	}

	nickName, ok := user["nickname"].(string)
	if !ok {
		nickName = ""
	}

	exp, ok := tokenClaims["exp"].(int)
	if !ok {
		exp = 0
	}

	userClaims := UserClaims{
		Avatar:        avatar,
		Email:         email,
		EmailVerified: emailVerified,
		Name:          name,
		NickName:      nickName,
	}

	return Claims{
		User: userClaims,
		Exp:  exp,
	}
}
