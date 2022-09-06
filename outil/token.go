package outil

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

const (
	jwtKey = "pfomoy#274/vp,"
)

func TokenGenWithClaims(claims map[string]any, method jwt.SigningMethod, key any) (string, error) {
	if key == nil {
		key = jwtKey
	}
	return jwt.NewWithClaims(method, jwt.MapClaims(claims)).SignedString(key)
}

func TokenParse(tokenStr string, key any) (claims map[string]any, err error) {
	if key == nil {
		key = jwtKey
	}
	token, err := jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("解析结果无法转换为jwt.MapClaims")
		return
	}
	return claims, nil
}

func TokenParseField(tokenStr string, key any, field string) (value any, err error) {
	if key == nil {
		key = jwtKey
	}
	token, err := jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("解析结果无法转换为jwt.MapClaims")
		return
	}
	return claims[field], nil
}
