package auth

import (
	"api/src/config"
	models_user "api/src/models/user"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func AuthTokenGen(user models_user.User) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["userID"] = user.ID
	permissions["exp"] = time.Now().Add(time.Hour * 4).UnixMilli()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.JwtSecretKey))
}

func AuthTokenValidate(r *http.Request) error {
	token, err := parseToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")
}

// Helpers
func authTokenExtractStr(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func AuthTokenExtractDataUser(r *http.Request) (uint, error) {
	token, err := parseToken(r)
	if err != nil {
		return 0, err
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return uint(userID), nil
	}
	return 0, errors.New("invalid token")
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	tokenStr := authTokenExtractStr(r)
	return jwt.Parse(tokenStr, keyChecker)
}

func keyChecker(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("signing method invalid: %v", token.Header["alg"])
	}
	return config.JwtSecretKey, nil
}

// A function AuthTokenExtractExpTime that return the timestamp of the token expiration time. The parameter is the token string.
func AuthTokenExtractExpTime(tokenStr string) (int64, error) {
	token, err := jwt.Parse(tokenStr, keyChecker)
	if err != nil {
		return 0, err
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int64(permissions["exp"].(float64)), nil
	}
	return 0, errors.New("invalid token")
}
