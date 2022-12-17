package utils

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func encodeToken(payload map[string]interface{}) (string, error) {
	_, tokenString, err := TokenAuth.Encode(payload)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewToken(userID uint) (string, error) {
	userIdMap := map[string]interface{}{"id": userID}
	tokenString, err := encodeToken(userIdMap)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(r *http.Request) (map[string]interface{}, error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return nil, err
	}

	return claims, nil
}
