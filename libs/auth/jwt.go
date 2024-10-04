package auth

import (
	"encoding/base64"

	"github.com/AcordoCertoBR/cp-atende-api/libs/errors"
	"github.com/golang-jwt/jwt"
)

func ValidateJWT(jwtSecret string, receivedToken string, expectedTokenType string) (retVal bool, err error) {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		secretKey, err := base64.StdEncoding.DecodeString(jwtSecret)
		if err != nil {
			return "", errors.Wrap(err)
		}
		return secretKey, nil
	})

	if err != nil {
		return retVal, errors.Wrap(err)
	}

	retVal = token.Valid

	// Token Information
	//_, ok := token.Claims.(jwt.MapClaims)
	//if !ok || !token.Valid {
	//	return retVal, errors.Wrap(err)
	//}

	return retVal, nil
}
