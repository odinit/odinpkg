package odinutil

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	JwtErrNotMapClaims  = errors.New("token.Claims不是MapClaims")
	JwtErrTokenNotValid = errors.New("token.Valid是false")
)

const (
	JwtKeyDefault = "pfomoy#274/vp,"
)

var (
	JwtMethodDefault = jwt.SigningMethodHS256
)

func JwtNewWithMapClaims(fields map[string]any, method jwt.SigningMethod) (string, error) {
	return jwt.NewWithClaims(method, jwt.MapClaims(fields)).SigningString()
}

func JwtNewWithMapClaimsHS256(fields map[string]any) (tokenStr string, err error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(fields)).SigningString()
}

func JwtParseMapClaims(tokenStr string) (tokenMap map[string]any, err error) {
	return JwtParseMapClaimsFunc(tokenStr, func(*jwt.Token) (interface{}, error) {
		return JwtKeyDefault, nil
	})
}

func JwtParseMapClaimsFunc(tokenStr string, keyFunc jwt.Keyfunc) (tokenMap map[string]any, err error) {
	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid {
		return nil, JwtErrTokenNotValid
	}
	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, JwtErrNotMapClaims
	}
	return tokenClaims, nil
}
